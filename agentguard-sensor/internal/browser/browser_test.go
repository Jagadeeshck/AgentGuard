package browser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseManifest(t *testing.T) {
	d := t.TempDir()
	p := filepath.Join(d, "manifest.json")
	_ = os.WriteFile(p, []byte(`{"name":"AI helper","permissions":["tabs","<all_urls>"]}`), 0o644)
	e, ok := ParseManifest(p)
	if !ok || !e.Risky {
		t.Fatal("expected risky extension")
	}
}
