package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	c, err := Load("../../testdata/config.yml")
	if err != nil {
		t.Fatal(err)
	}
	if c.Output.Format != "ndjson" {
		t.Fatal("bad format")
	}
}
func TestDefaults(t *testing.T) {
	c := DefaultConfig()
	if c.Watch.Interval <= 0 || c.Output.Path == "" {
		t.Fatal("missing defaults")
	}
}
func TestEnsureRuntimePaths(t *testing.T) {
	d := t.TempDir()
	c := DefaultConfig()
	c.Output.Path = filepath.Join(d, "o", "findings.ndjson")
	c.Watch.StateFile = filepath.Join(d, "s", "state.json")
	if err := EnsureRuntimePaths(c); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Dir(c.Output.Path)); err != nil {
		t.Fatal(err)
	}
}
func TestPrivacyGuardrails(t *testing.T) {
	c := DefaultConfig()
	c.Privacy.CollectClipboard = true
	if c.Validate() == nil {
		t.Fatal("expected error")
	}
}
