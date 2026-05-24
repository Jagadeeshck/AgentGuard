package findings

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type Event struct {
	Timestamp  string         `json:"@timestamp"`
	ECS        ECS            `json:"ecs"`
	Event      EventMeta      `json:"event"`
	Observer   Observer       `json:"observer"`
	Host       Host           `json:"host"`
	AgentGuard AISentinelData `json:"agentguard"`
	AIS        AISentinelData `json:"ai_sentinel,omitempty"`
}

type ECS struct {
	Version string `json:"version"`
}
type EventMeta struct {
	Module    string   `json:"module"`
	Dataset   string   `json:"dataset"`
	Kind      string   `json:"kind"`
	Category  []string `json:"category"`
	Type      []string `json:"type"`
	Action    string   `json:"action"`
	Outcome   string   `json:"outcome"`
	RiskScore int      `json:"risk_score"`
	Severity  int      `json:"severity"`
}
type Observer struct {
	Vendor  string `json:"vendor"`
	Product string `json:"product"`
	Type    string `json:"type"`
}
type Host struct {
	Name string `json:"name"`
}
type AISentinelData struct {
	Schema  SchemaVersion  `json:"schema"`
	Finding Finding        `json:"finding"`
	Risk    Risk           `json:"risk"`
	Allowed bool           `json:"allowed"`
	Details map[string]any `json:"details,omitempty"`
}
type SchemaVersion struct {
	Version string `json:"version"`
}
type Finding struct {
	ID         string  `json:"id"`
	Type       string  `json:"type"`
	Name       string  `json:"name"`
	Status     string  `json:"status"`
	Confidence float64 `json:"confidence"`
}
type Risk struct {
	Level   string   `json:"level"`
	Score   int      `json:"score"`
	Reasons []string `json:"reasons"`
}

func NewEvent(ftype, id, name string, score int, reasons []string, details map[string]any) Event {
	h, _ := os.Hostname()
	data := AISentinelData{Schema: SchemaVersion{Version: "1.0.0-draft"}, Finding: Finding{ID: id, Type: ftype, Name: name, Status: "open", Confidence: 0.7}, Risk: Risk{Level: ScoreToLevel(score), Score: score, Reasons: reasons}, Details: details}
	return Event{Timestamp: time.Now().UTC().Format(time.RFC3339), ECS: ECS{Version: "8.11.0"},
		Event:      EventMeta{Module: "agentguard", Dataset: "agentguard.findings", Kind: "alert", Category: []string{"configuration", "process"}, Type: []string{"info"}, Action: "finding_detected", Outcome: "success", RiskScore: score, Severity: score},
		Observer:   Observer{Vendor: "AgentGuard", Product: "AgentGuard Sensor", Type: "endpoint"},
		Host:       Host{Name: h},
		AgentGuard: data,
		AIS:        data,
	}
}

func StableFindingID(prefix string, parts ...string) string {
	h := sha256.Sum256([]byte(strings.Join(parts, "|")))
	return fmt.Sprintf("%s%s", prefix, hex.EncodeToString(h[:])[:16])
}

func SafeFingerprint(e Event) string {
	details := e.AgentGuard.Details
	if len(details) == 0 {
		details = e.AIS.Details
	}
	keys := make([]string, 0, len(details))
	for k := range details {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	findingType := e.AgentGuard.Finding.Type
	riskScore := e.AgentGuard.Risk.Score
	if findingType == "" {
		findingType = e.AIS.Finding.Type
	}
	if riskScore == 0 {
		riskScore = e.AIS.Risk.Score
	}
	chunks := []string{findingType, fmt.Sprintf("%d", riskScore)}
	for _, k := range keys {
		chunks = append(chunks, fmt.Sprintf("%s=%v", k, details[k]))
	}
	h := sha256.Sum256([]byte(strings.Join(chunks, "|")))
	return hex.EncodeToString(h[:])
}

func ScoreToLevel(s int) string {
	if s >= 86 {
		return "critical"
	}
	if s >= 61 {
		return "high"
	}
	if s >= 31 {
		return "medium"
	}
	return "low"
}
func ValidateRequired(e Event) error {
	findingID := e.AgentGuard.Finding.ID
	if findingID == "" {
		findingID = e.AIS.Finding.ID
	}
	if e.Event.Module == "" || findingID == "" || e.Timestamp == "" {
		return fmt.Errorf("missing required fields")
	}
	return nil
}
