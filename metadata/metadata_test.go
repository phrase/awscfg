package metadata

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestInstanceID(t *testing.T) {
	defer setupMockEndpoint()()

	id, err := InstanceID()
	if err != nil {
		t.Fatal(err)
	}
	if v, ex := id, "i-123456"; v != ex {
		t.Errorf("expected instance id to be %q, was %q", ex, v)
	}
}

func TestAvailabilityZone(t *testing.T) {
	defer setupMockEndpoint()()
	zone, err := AvailabilityZone()
	if err != nil {
		t.Fatal(err)
	}
	if v, ex := zone, "eu-central-1a"; v != ex {
		t.Errorf("expected availability zone to be %q, was %q", ex, v)
	}
}

func TestIAMRoles(t *testing.T) {
	defer setupMockEndpoint()()

	roles, err := IAMRoles()
	if err != nil {
		t.Fatal(err)
	}
	if v, ex := strings.Join(roles, ","), "role1,role2"; v != ex {
		t.Errorf("expected roles to be %q, was %q", ex, v)
	}
}

func TestIAMCredentials(t *testing.T) {
	defer setupMockEndpoint()()
	c, err := IAMCredentials("role1")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		Name     string
		Expected interface{}
		Value    interface{}
	}{
		{"AccessKeyId", "ACCESS_KEY_ID", c.AccessKeyId},
		{"SecretAccessKey", "SECRET_ACCESS_KEY", c.SecretAccessKey},
		{"Token", "TOKEN", c.Token},
	}

	for _, tst := range tests {
		if tst.Expected != tst.Value {
			t.Errorf("expected %s to be %#v, was %#v", tst.Name, tst.Expected, tst.Value)
		}
	}
}

func setupMockEndpoint() func() {
	ep := endpoint
	endpoint = mockEndpoint().URL
	return func() {
		endpoint = ep
	}
}

func mockEndpoint() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/latest/meta-data/instance-id", serveText("i-123456"))
	mux.HandleFunc("/latest/meta-data/iam/security-credentials/", serveText("role1\nrole2"))
	mux.HandleFunc("/latest/meta-data/iam/security-credentials/role1", serveText(sg))
	mux.HandleFunc("/latest/meta-data/placement/availability-zone/", serveText("eu-central-1a"))
	return httptest.NewServer(mux)
}

func serveText(txt string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, txt)
	}
}

const sg = `{
  "Code" : "Success",
  "LastUpdated" : "2015-10-21T05:01:21Z",
  "Type" : "AWS-HMAC",
  "AccessKeyId" : "ACCESS_KEY_ID",
  "SecretAccessKey" : "SECRET_ACCESS_KEY",
  "Token" : "TOKEN",
  "Expiration" : "2015-10-21T11:02:30Z"
}`
