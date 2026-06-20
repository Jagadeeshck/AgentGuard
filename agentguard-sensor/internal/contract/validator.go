package contract

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var prohibitedFields = []string{
	"agentguard.prompt", "agentguard.prompts", "agentguard.completion", "agentguard.response", "browser.history", "http.request.body.content",
}

func ValidateEventMap(event map[string]any) error {
	required := []string{"@timestamp", "event.module", "event.dataset", "event.kind", "event.category", "event.type", "event.action", "event.outcome", "observer.vendor", "observer.product", "host.name", "agentguard.schema.version", "agentguard.finding.id", "agentguard.finding.type", "agentguard.risk.level", "agentguard.risk.score"}
	for _, field := range required {
		if _, ok := getPath(event, field); !ok {
			return fmt.Errorf("missing required field: %s", field)
		}
	}
	if _, err := time.Parse(time.RFC3339, fmt.Sprintf("%v", event["@timestamp"])); err != nil {
		return fmt.Errorf("@timestamp must be RFC3339: %w", err)
	}
	if mod, _ := getPath(event, "event.module"); mod != "agentguard" {
		return fmt.Errorf("event.module must be agentguard")
	}
	if ds, _ := getPath(event, "event.dataset"); ds != "agentguard.findings" {
		return fmt.Errorf("event.dataset must be agentguard.findings")
	}
	if ver, _ := getPath(event, "agentguard.schema.version"); ver != "1.0.0-draft" {
		return fmt.Errorf("unsupported schema version: %v", ver)
	}

	if category, _ := getPath(event, "event.category"); !nonEmptyStringArray(category) {
		return fmt.Errorf("event.category must be a non-empty string array")
	}
	if typ, _ := getPath(event, "event.type"); !nonEmptyStringArray(typ) {
		return fmt.Errorf("event.type must be a non-empty string array")
	}
	if score, _ := getPath(event, "agentguard.risk.score"); !numberInRange(score, 0, 100) {
		return fmt.Errorf("agentguard.risk.score must be a number from 0 to 100")
	}
	for _, p := range prohibitedFields {
		if _, ok := getPath(event, p); ok {
			return fmt.Errorf("prohibited field present: %s", p)
		}
	}
	return nil
}

func ValidateNDJSONFile(path string) (int, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	for i, line := range strings.Split(strings.TrimSpace(string(b)), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		var e map[string]any
		if err := json.Unmarshal([]byte(line), &e); err != nil {
			return 0, fmt.Errorf("line %d invalid json: %w", i+1, err)
		}
		if err := ValidateEventMap(e); err != nil {
			return 0, fmt.Errorf("line %d invalid contract: %w", i+1, err)
		}
	}
	count := 0
	for _, line := range strings.Split(strings.TrimSpace(string(b)), "\n") {
		if strings.TrimSpace(line) != "" {
			count++
		}
	}
	if count == 0 {
		return 0, fmt.Errorf("no events found")
	}
	return count, nil
}

func RepoRootFromWD() string {
	wd, _ := os.Getwd()
	return filepath.Clean(filepath.Join(wd, ".."))
}

func getPath(m map[string]any, path string) (any, bool) {
	parts := strings.Split(path, ".")
	var cur any = m
	for _, p := range parts {
		mm, ok := cur.(map[string]any)
		if !ok {
			return nil, false
		}
		cur, ok = mm[p]
		if !ok {
			return nil, false
		}
	}
	return cur, true
}

func nonEmptyStringArray(value any) bool {
	items, ok := value.([]any)
	if !ok || len(items) == 0 {
		return false
	}
	for _, item := range items {
		if s, ok := item.(string); !ok || s == "" {
			return false
		}
	}
	return true
}

func numberInRange(value any, min, max float64) bool {
	var n float64
	switch v := value.(type) {
	case float64:
		n = v
	case int:
		n = float64(v)
	default:
		return false
	}
	return n >= min && n <= max
}
