package contract

import (
	"encoding/json"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/agentguard/agentguard-sensor/internal/findings"
)

func TestSampleFixturesValidate(t *testing.T) {
	_, thisFile, _, _ := runtime.Caller(0)
	path := filepath.Clean(filepath.Join(filepath.Dir(thisFile), "..", "..", "..", "examples", "sample-findings", "v1-sample-events.ndjson"))
	if err := ValidateNDJSONFile(path); err != nil {
		t.Fatalf("fixtures failed validation: %v", err)
	}
}

func TestSensorEventShapeValidates(t *testing.T) {
	e := findings.NewEvent("agentguard.process.detected", "ag-test-1", "untrusted_agent", 79, []string{"unexpected parent process"}, map[string]any{"process_name": "agent-runner"})
	b, _ := json.Marshal(e)
	var m map[string]any
	_ = json.Unmarshal(b, &m)
	if err := ValidateEventMap(m); err != nil {
		t.Fatalf("sensor event does not satisfy contract: %v", err)
	}
}

func TestProhibitedFieldRejected(t *testing.T) {
	e := findings.NewEvent("agentguard.finding.detected", "ag-test-2", "x", 10, nil, map[string]any{})
	b, _ := json.Marshal(e)
	var m map[string]any
	_ = json.Unmarshal(b, &m)
	ag := m["agentguard"].(map[string]any)
	ag["prompt"] = "secret prompt text"
	if err := ValidateEventMap(m); err == nil {
		t.Fatal("expected prohibited prompt field to be rejected")
	}
}
