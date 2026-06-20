package startup

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanPaths(t *testing.T) {
	dir := t.TempDir()
	match := filepath.Join(dir, "cursor-agent.desktop")
	other := filepath.Join(dir, "ordinary.desktop")
	if err := os.WriteFile(match, []byte("Exec=cursor"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(other, []byte("Exec=ordinary"), 0o644); err != nil {
		t.Fatal(err)
	}
	got := scanPaths([]string{dir, filepath.Join(dir, "missing")})
	if len(got) != 1 {
		t.Fatalf("got %#v", got)
	}
	if got[0].Path != match || got[0].Command != match || !got[0].Match {
		t.Fatalf("unexpected item: %#v", got[0])
	}
}
