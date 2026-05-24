package findings

import (
	"fmt"
	"strings"
)

var prohibitedFieldPaths = []string{
	"agentguard.prompt",
	"ai_sentinel.prompt",
	"ai_sentinel.prompts",
	"ai_sentinel.prompt_content",
	"ai_sentinel.completion",
	"ai_sentinel.response",
	"ai_sentinel.secret",
	"ai_sentinel.secrets",
	"clipboard",
	"browser.history",
	"http.request.body.content",
}

func SyntheticTestEvents() []Event {
	timestamp := "2026-01-01T00:00:00Z"
	host := "synthetic-host-01"
	base := []struct {
		ftype   string
		name    string
		score   int
		reasons []string
		details map[string]any
	}{
		{"mcp_server", "mcp-gateway-alpha", 75, []string{"shell capability enabled", "filesystem write allowed"}, map[string]any{"command": "mcp-server", "args": []string{"--transport", "stdio"}, "capabilities": []string{"shell", "filesystem"}, "config_path": "/synthetic/etc/mcp/config.yaml"}},
		{"browser_extension", "agentguard-helper", 52, []string{"risky permission set"}, map[string]any{"id": "ag_ext_001", "name": "AgentGuard Helper", "path": "/synthetic/browser/extensions/ag_ext_001"}},
		{"local_llm_service", "local-llm-11434", 68, []string{"service exposed on all interfaces"}, map[string]any{"port": 11434, "bind": "0.0.0.0"}},
		{"startup_item", "agentguard-autostart", 43, []string{"untrusted startup path"}, map[string]any{"path": "/synthetic/startup/agentguard-autostart.sh"}},
		{"suspicious_agent_process", "agentguard-shadow", 88, []string{"binary path mismatch", "unexpected parent process"}, map[string]any{"process": map[string]any{"name": "agentguard-shadow", "pid": 4242, "ppid": 100, "executable": "/synthetic/bin/agentguard-shadow", "command_line": "agentguard-shadow --token [REDACTED]"}}},
	}

	events := make([]Event, 0, len(base))
	for _, b := range base {
		e := NewEvent(b.ftype, StableFindingID("ag-synth-", b.ftype, b.name), b.name, b.score, b.reasons, b.details)
		e.Timestamp = timestamp
		e.Host.Name = host
		e.Observer = Observer{Vendor: "AgentGuard", Product: "AgentGuard Sensor", Type: "endpoint"}
		events = append(events, e)
	}
	return events
}

func ValidateContractMap(event map[string]any) error {
	required := []string{
		"@timestamp", "ecs.version", "event.module", "event.dataset", "event.kind", "event.category", "event.type", "event.action", "event.outcome", "host.name", "observer.vendor", "observer.product", "observer.type", "agentguard.schema.version", "agentguard.finding.id", "agentguard.finding.type", "agentguard.risk.score",
	}
	for _, p := range required {
		if !hasPath(event, p) {
			return fmt.Errorf("missing required field: %s", p)
		}
	}
	if v, ok := getPath(event, "event.module"); !ok || v != "agentguard" {
		return fmt.Errorf("event.module must be agentguard")
	}
	if v, ok := getPath(event, "event.dataset"); !ok || v != "agentguard.findings" {
		return fmt.Errorf("event.dataset must be agentguard.findings")
	}
	for _, p := range prohibitedFieldPaths {
		if hasPath(event, p) {
			return fmt.Errorf("prohibited field present: %s", p)
		}
	}
	return nil
}

func hasPath(m map[string]any, path string) bool { _, ok := getPath(m, path); return ok }
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
