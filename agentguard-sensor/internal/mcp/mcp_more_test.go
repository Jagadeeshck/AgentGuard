package mcp

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParseFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "mcp.json")
	json := `{"mcpServers":{"fs":{"command":"npx","args":["filesystem-server","--root","/tmp",7]}}}`
	if err := os.WriteFile(path, []byte(json), 0o644); err != nil {
		t.Fatal(err)
	}
	got, ok := parseFile(path)
	if !ok || len(got) != 1 {
		t.Fatalf("ok=%v got=%#v", ok, got)
	}
	if got[0].Name != "fs" || got[0].Command != "npx" || !reflect.DeepEqual(got[0].Args, []string{"filesystem-server", "--root", "/tmp"}) {
		t.Fatalf("unexpected server %#v", got[0])
	}
	if len(got[0].Capabilities) == 0 {
		t.Fatal("expected capabilities")
	}
}

func TestParseFileInvalidAndMissing(t *testing.T) {
	if got, ok := parseFile(filepath.Join(t.TempDir(), "missing.json")); ok || got != nil {
		t.Fatalf("missing ok=%v got=%#v", ok, got)
	}
	path := filepath.Join(t.TempDir(), "bad.json")
	if err := os.WriteFile(path, []byte("{"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got, ok := parseFile(path); ok || got != nil {
		t.Fatalf("bad ok=%v got=%#v", ok, got)
	}
}
