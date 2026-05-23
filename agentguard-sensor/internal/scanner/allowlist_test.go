package scanner

import "testing"

func TestAllowlistMatch(t *testing.T) {
	if !contains([]string{"a"}, "a") {
		t.Fatal("expected match")
	}
}
