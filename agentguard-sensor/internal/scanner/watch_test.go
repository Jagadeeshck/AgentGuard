package scanner

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/agentguard/agentguard-sensor/internal/findings"
	"github.com/agentguard/agentguard-sensor/internal/output"
)

func TestStableFindingID(t *testing.T) {
	a := findings.StableFindingID("ag-mcp-", "a", "b")
	b := findings.StableFindingID("ag-mcp-", "a", "b")
	if a != b {
		t.Fatal("ids must be stable")
	}
}

func TestStateSaveLoad(t *testing.T) {
	p := filepath.Join(t.TempDir(), "state.json")
	s := watchState{Findings: map[string]stateRecord{"id": {ID: "id", Type: "mcp_server", Fingerprint: "x"}}}
	if err := saveState(p, s); err != nil {
		t.Fatal(err)
	}
	if got := loadState(p); got.Findings["id"].Fingerprint != "x" {
		t.Fatal("load failed")
	}
}

func TestIntervalParsing(t *testing.T) {
	d, err := time.ParseDuration("30s")
	if err != nil || d != 30*time.Second {
		t.Fatal("interval parse")
	}
}

func TestOutputAppendMode(t *testing.T) {
	p := filepath.Join(t.TempDir(), "f.ndjson")
	e := findings.NewEvent("mcp_server", "id", "n", 10, nil, nil)
	_ = output.WriteAppend([]findings.Event{e}, p, false, false)
	_ = output.WriteAppend([]findings.Event{e}, p, false, false)
	b, _ := os.ReadFile(p)
	if len(b) == 0 {
		t.Fatal("expected data")
	}
}

func TestAllowlist(t *testing.T) {
	al := loadAllowlist(filepath.Join("testdata", "allowlist.yml"))
	if len(al.FindingIDs) == 0 {
		t.Fatal("allowlist not loaded")
	}
}
