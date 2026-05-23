package mcp

import "testing"

func TestInferCapabilities(t *testing.T) {
	c := infer("npx", []string{"mcp-server", "--root", "/tmp"})
	if len(c) == 0 {
		t.Fatal("expected capabilities")
	}
}
