package awscfg

import "testing"

func TestValidCode(t *testing.T) {
	valid := []string{"000000", "000001", "100000", "123456"}
	for _, c := range valid {
		if !validCode(c) {
			t.Errorf("expected code %q to be valid", c)
		}
	}

	notValid := []string{"00000", "00000a", " 00000"}

	for _, c := range notValid {
		if validCode(c) {
			t.Errorf("expected code %q not to be valid", c)
		}
	}
}
