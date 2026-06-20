package scanner

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/agentguard/agentguard-sensor/internal/browser"
	"github.com/agentguard/agentguard-sensor/internal/findings"
	"github.com/agentguard/agentguard-sensor/internal/localai"
	"github.com/agentguard/agentguard-sensor/internal/mcp"
	"github.com/agentguard/agentguard-sensor/internal/output"
	"github.com/agentguard/agentguard-sensor/internal/process"
	"github.com/agentguard/agentguard-sensor/internal/rules"
	"github.com/agentguard/agentguard-sensor/internal/startup"
)

type Options struct {
	OutputPath     string
	Stdout, Pretty bool
	AllowlistPath  string
}
type Allowlist struct {
	ProcessPaths, MCPConfigPaths, BrowserExtensionIDs, FindingIDs []string
	LocalPorts                                                    []int
}

func Run(opts Options) error {
	events, _ := ScanFindings(opts.AllowlistPath)
	return output.Write(events, opts.OutputPath, opts.Stdout, opts.Pretty)
}

func ScanFindings(allowlistPath string) ([]findings.Event, Allowlist) {
	al := loadAllowlist(allowlistPath)
	events := []findings.Event{}
	for _, s := range mcp.Scan() {
		score, reasons := rules.ComputeScore(rules.ScoreInput{MCPShell: has(s.Capabilities, "shell"), MCPFilesystem: has(s.Capabilities, "filesystem")})
		id := findings.StableFindingID("ag-mcp-", s.ConfigPath, s.Name, s.Command)
		e := findings.NewEvent("mcp_server", id, s.Name, score, reasons, map[string]any{"command": s.Command, "args": s.Args, "capabilities": s.Capabilities, "config_path": s.ConfigPath})
		if contains(al.MCPConfigPaths, s.ConfigPath) || contains(al.FindingIDs, id) {
			e.AgentGuard.Allowed = true
			e.AgentGuard.Finding.Status = "allowed"
			e.Event.Action = "finding_allowed"
		}
		events = append(events, e)
	}
	for _, p := range process.Scan() {
		l := strings.ToLower(p.Name + " " + p.CommandLine)
		if hasAny(l, []string{"ollama", "lmstudio", "claude", "cursor", "open-interpreter", "autogpt", "crewai", "langchain", "llamaindex", "copilot"}) {
			red, rf := rules.RedactSecrets(p.CommandLine)
			score, reasons := rules.ComputeScore(rules.ScoreInput{UnknownRuntimeAI: hasAny(l, []string{"python", "node", "npx", "bun", "deno", "uv"}), RedactedSecret: rf, SecurityKeywords: hasAny(l, []string{"security", "cyber"})})
			id := findings.StableFindingID("ag-proc-", p.Executable, red)
			e := findings.NewEvent("suspicious_agent_process", id, p.Name, score, reasons, map[string]any{"process": map[string]any{"name": p.Name, "pid": p.PID, "ppid": p.PPID, "executable": p.Executable, "command_line": red}})
			if contains(al.ProcessPaths, p.Executable) || contains(al.FindingIDs, id) {
				e.AgentGuard.Allowed = true
				e.AgentGuard.Finding.Status = "allowed"
				e.Event.Action = "finding_allowed"
			}
			events = append(events, e)
		}
	}
	for _, s := range localai.Scan() {
		score, reasons := rules.ComputeScore(rules.ScoreInput{LocalExposed: s.Exposed})
		id := findings.StableFindingID("ag-llm-", s.Addr, fmt.Sprintf("%d", s.Port), s.ProcessName)
		e := findings.NewEvent("local_llm_service", id, fmt.Sprintf("port_%d", s.Port), score, reasons, map[string]any{"port": s.Port, "bind": s.Addr, "process_name": s.ProcessName})
		if containsInt(al.LocalPorts, s.Port) || contains(al.FindingIDs, id) {
			e.AgentGuard.Allowed = true
			e.AgentGuard.Finding.Status = "allowed"
			e.Event.Action = "finding_allowed"
		}
		events = append(events, e)
	}
	for _, it := range startup.Scan() {
		score, reasons := rules.ComputeScore(rules.ScoreInput{Startup: true})
		id := findings.StableFindingID("ag-startup-", it.Path, it.Command)
		e := findings.NewEvent("startup_item", id, it.Path, score, reasons, map[string]any{"path": it.Path, "command": it.Command})
		events = append(events, e)
	}
	for _, ext := range browser.Scan() {
		score, reasons := rules.ComputeScore(rules.ScoreInput{})
		id := findings.StableFindingID("ag-ext-", ext.Browser, ext.Profile, ext.ID)
		e := findings.NewEvent("browser_extension", id, ext.Name, score, reasons, map[string]any{"browser": ext.Browser, "profile": ext.Profile, "extension_id": ext.ID, "path": ext.Path})
		if contains(al.BrowserExtensionIDs, ext.ID) || contains(al.FindingIDs, id) {
			e.AgentGuard.Allowed = true
			e.AgentGuard.Finding.Status = "allowed"
			e.Event.Action = "finding_allowed"
		}
		events = append(events, e)
	}
	return events, al
}

func List(kind string) ([]string, error) { return []string{"not implemented for MVP"}, nil }
func has(a []string, k string) bool {
	for _, x := range a {
		if x == k {
			return true
		}
	}
	return false
}
func hasAny(s string, ks []string) bool {
	for _, k := range ks {
		if strings.Contains(s, k) {
			return true
		}
	}
	return false
}
func contains(a []string, v string) bool {
	for _, x := range a {
		if x == v {
			return true
		}
	}
	return false
}
func containsInt(a []int, v int) bool {
	for _, x := range a {
		if x == v {
			return true
		}
	}
	return false
}

func loadAllowlist(path string) Allowlist {
	if path == "" {
		return Allowlist{}
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return Allowlist{}
	}
	lines := strings.Split(string(b), "\n")
	a := Allowlist{}
	cur := ""
	for _, ln := range lines {
		ln = strings.TrimSpace(ln)
		if ln == "" || strings.HasPrefix(ln, "#") {
			continue
		}
		if strings.HasSuffix(ln, ":") {
			cur = strings.TrimSuffix(ln, ":")
			continue
		}
		if strings.Contains(ln, ":") && !strings.HasPrefix(ln, "-") {
			parts := strings.SplitN(ln, ":", 2)
			cur = strings.TrimSpace(parts[0])
			for _, v := range inlineListValues(parts[1]) {
				applyAllowlistValue(&a, cur, v)
			}
			continue
		}
		if strings.HasPrefix(ln, "-") {
			v := strings.TrimSpace(strings.TrimPrefix(ln, "-"))
			applyAllowlistValue(&a, cur, v)
		}
	}
	return a
}

func applyAllowlistValue(a *Allowlist, key, value string) {
	value = strings.Trim(strings.TrimSpace(value), `"'`)
	if value == "" {
		return
	}
	switch key {
	case "process_paths":
		a.ProcessPaths = append(a.ProcessPaths, value)
	case "mcp_config_paths":
		a.MCPConfigPaths = append(a.MCPConfigPaths, value)
	case "browser_extension_ids":
		a.BrowserExtensionIDs = append(a.BrowserExtensionIDs, value)
	case "finding_ids":
		a.FindingIDs = append(a.FindingIDs, value)
	case "local_ports":
		port, err := strconv.Atoi(value)
		if err == nil {
			a.LocalPorts = append(a.LocalPorts, port)
		}
	}
}

func inlineListValues(raw string) []string {
	raw = strings.TrimSpace(raw)
	if !strings.HasPrefix(raw, "[") || !strings.HasSuffix(raw, "]") {
		return nil
	}
	raw = strings.TrimSuffix(strings.TrimPrefix(raw, "["), "]")
	if strings.TrimSpace(raw) == "" {
		return nil
	}
	parts := strings.Split(raw, ",")
	values := make([]string, 0, len(parts))
	for _, part := range parts {
		values = append(values, strings.TrimSpace(part))
	}
	return values
}
