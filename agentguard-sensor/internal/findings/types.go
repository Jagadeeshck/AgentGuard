package findings

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"
)

type Event struct {
	Timestamp string         `json:"@timestamp"`
	ECS       ECS            `json:"ecs"`
	Event     EventMeta      `json:"event"`
	Observer  Observer       `json:"observer"`
	Host      Host           `json:"host"`
	AIS       AISentinelData `json:"ai_sentinel"`
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
	Finding Finding        `json:"finding"`
	Risk    Risk           `json:"risk"`
	Allowed bool           `json:"allowed"`
	Details map[string]any `json:"details,omitempty"`
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

func NewEvent(ftype, name string, score int, reasons []string, details map[string]any) Event {
	h, _ := os.Hostname()
	level := ScoreToLevel(score)
	id := hashID(ftype + name + strings.Join(reasons, "|"))
	return Event{Timestamp: time.Now().UTC().Format(time.RFC3339), ECS: ECS{Version: "8.11.0"},
		Event:    EventMeta{Module: "ai_sentinel", Dataset: "ai_sentinel.findings", Kind: "alert", Category: []string{"configuration", "process"}, Type: []string{"info"}, Action: ftype, Outcome: "success", RiskScore: score, Severity: score},
		Observer: Observer{Vendor: "AgentGuard", Product: "AgentGuard Sensor", Type: "endpoint"}, Host: Host{Name: h},
		AIS: AISentinelData{Finding: Finding{ID: id, Type: ftype, Name: name, Status: "open", Confidence: 0.7}, Risk: Risk{Level: level, Score: score, Reasons: reasons}, Allowed: false, Details: details},
	}
}

func hashID(in string) string { s := sha1.Sum([]byte(in)); return hex.EncodeToString(s[:8]) }
func ScoreToLevel(s int) string {
	switch {
	case s >= 86:
		return "critical"
	case s >= 61:
		return "high"
	case s >= 31:
		return "medium"
	default:
		return "low"
	}
}
func ValidateRequired(e Event) error {
	if e.Event.Module == "" || e.AIS.Finding.ID == "" || e.Timestamp == "" {
		return fmt.Errorf("missing required fields")
	}
	return nil
}
