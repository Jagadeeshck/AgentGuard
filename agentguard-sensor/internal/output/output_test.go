package output

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/agentguard/agentguard-sensor/internal/findings"
)

func sampleEvent() findings.Event {
	return findings.NewEvent("test_kind", "id1", "name1", 10, []string{"reason"}, map[string]any{"k": "v"})
}

func TestWriteAndWriteAppend(t *testing.T) {
	path := filepath.Join(t.TempDir(), "nested", "findings.ndjson")
	e := sampleEvent()
	if err := Write([]findings.Event{e}, path, false, false); err != nil {
		t.Fatal(err)
	}
	if err := WriteAppend([]findings.Event{e}, path, false, true); err != nil {
		t.Fatal(err)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	if len(lines) < 2 {
		t.Fatalf("expected appended lines, got %q", string(b))
	}
	var decoded findings.Event
	if err := json.Unmarshal([]byte(lines[0]), &decoded); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(b), "\n  ") {
		t.Fatalf("expected pretty JSON in append output: %q", string(b))
	}
}

func TestWriteToStdout(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	old := os.Stdout
	os.Stdout = w
	if err := Write([]findings.Event{sampleEvent()}, "", true, true); err != nil {
		t.Fatal(err)
	}
	_ = w.Close()
	os.Stdout = old
	b, err := io.ReadAll(r)
	_ = r.Close()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(b), "test_kind") || !strings.Contains(string(b), "\n  ") {
		t.Fatalf("unexpected stdout: %q", string(b))
	}
}
