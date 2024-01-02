package common

import "testing"

func TestRandomString(t *testing.T) {
	s := RandomString(32)

	if len(s) != 32 {
		t.Errorf("expected string with length 32, got %d", len(s))
	}
}
