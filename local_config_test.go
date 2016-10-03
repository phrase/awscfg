package awscfg

import (
	"strings"
	"testing"
)

func TestNewFromLocalConfig(t *testing.T) {
	m, err := parseLocalConfig(strings.NewReader(localConfigTpl))
	if err != nil {
		t.Fatal(err)
	}
	d, ok := m["default"]
	if !ok {
		t.Errorf("expected map to have key default")
	}
	if v, ex := d["aws_access_key_id"], "key"; ex != v {
		t.Errorf("expected value.AccessKeyID to be %q, was %q", ex, v)
	}
	if v, ex := d["aws_secret_access_key"], "secret"; ex != v {
		t.Errorf("expected value.SecretAccessKey to be %q, was %q", ex, v)
	}
}

const localConfigTpl = `[default]
aws_access_key_id = key
aws_secret_access_key = secret
`
