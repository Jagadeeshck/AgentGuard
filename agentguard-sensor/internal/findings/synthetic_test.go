package findings

import (
	"encoding/json"
	"testing"
)

func TestSyntheticTestEventsUseCanonicalContract(t *testing.T) {
	events := SyntheticTestEvents()
	if len(events) == 0 {
		t.Fatal("expected synthetic events")
	}
	for _, e := range events {
		if e.Event.Module != "agentguard" || e.Event.Dataset != "agentguard.findings" {
			t.Fatalf("event uses non-canonical dataset: %#v", e.Event)
		}
		if e.AgentGuard.Finding.ID == "" || e.AgentGuard.Risk.Score == 0 {
			t.Fatalf("event missing agentguard finding/risk: %#v", e.AgentGuard)
		}
		b, _ := json.Marshal(e)
		if !json.Valid(b) {
			t.Fatalf("event is not valid json: %s", b)
		}
	}
}
