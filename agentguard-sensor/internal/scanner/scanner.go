package scanner

import (
	"fmt"
	"os"
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
	ProcessPaths, MCPConfigPaths, BrowserExtensionIDs, FindingIDs []string `yaml:"process_paths"`
	LocalPorts                                                    []int    `yaml:"local_ports"`
}

func Run(opts Options) error {
	al := loadAllowlist(opts.AllowlistPath)
	events := []findings.Event{}
	for _, s := range mcp.Scan() {
		score, reasons := rules.ComputeScore(rules.ScoreInput{MCPShell: has(s.Capabilities, "shell"), MCPFilesystem: has(s.Capabilities, "filesystem")})
		e := findings.NewEvent("mcp_server", s.Name, score, reasons, map[string]any{"command": s.Command, "args": s.Args, "capabilities": s.Capabilities, "config_path": s.ConfigPath})
		if contains(al.MCPConfigPaths, s.ConfigPath) {
			e.AIS.Allowed = true
			e.AIS.Finding.Status = "allowed"
		}
		events = append(events, e)
	}
	for _, p := range process.Scan() {
		l := strings.ToLower(p.Name + " " + p.CommandLine)
		if hasAny(l, []string{"ollama", "lmstudio", "claude", "cursor", "open-interpreter", "autogpt", "crewai", "langchain", "llamaindex", "copilot"}) {
			red, rf := rules.RedactSecrets(p.CommandLine)
			score, reasons := rules.ComputeScore(rules.ScoreInput{UnknownRuntimeAI: hasAny(l, []string{"python", "node", "npx", "bun", "deno", "uv"}), RedactedSecret: rf, SecurityKeywords: hasAny(l, []string{"security", "cyber"})})
			e := findings.NewEvent("suspicious_agent_process", p.Name, score, reasons, map[string]any{"process": map[string]any{"name": p.Name, "pid": p.PID, "ppid": p.PPID, "executable": p.Executable, "command_line": red}})
			events = append(events, e)
		}
	}
	for _, s := range localai.Scan() {
		score, reasons := rules.ComputeScore(rules.ScoreInput{LocalExposed: s.Exposed})
		e := findings.NewEvent("local_llm_service", fmt.Sprintf("port_%d", s.Port), score, reasons, map[string]any{"port": s.Port, "bind": s.Addr})
		if containsInt(al.LocalPorts, s.Port) {
			e.AIS.Allowed = true
			e.AIS.Finding.Status = "allowed"
		}
		events = append(events, e)
	}
	for _, it := range startup.Scan() {
		score, reasons := rules.ComputeScore(rules.ScoreInput{Startup: true})
		events = append(events, findings.NewEvent("startup_item", it.Path, score, reasons, map[string]any{"path": it.Path}))
	}
	_ = browser.ParseManifest
	return output.Write(events, opts.OutputPath, opts.Stdout, opts.Pretty)
}
func List(kind string) ([]string, error) {
	switch kind {
	case "list-mcp":
		r := []string{}
		for _, s := range mcp.Scan() {
			r = append(r, s.Name+" "+s.Command)
		}
		return r, nil
	case "list-local-ai":
		r := []string{}
		for _, s := range localai.Scan() {
			r = append(r, fmt.Sprintf("%d %s", s.Port, s.Addr))
		}
		return r, nil
	case "list-startup":
		r := []string{}
		for _, s := range startup.Scan() {
			r = append(r, s.Path)
		}
		return r, nil
	case "list-processes":
		r := []string{}
		for _, p := range process.Scan() {
			r = append(r, p.Name)
		}
		return r, nil
	default:
		return []string{"not implemented for MVP"}, nil
	}
}
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
		if strings.HasPrefix(ln, "-") {
			v := strings.TrimSpace(strings.TrimPrefix(ln, "-"))
			switch cur {
			case "process_paths":
				a.ProcessPaths = append(a.ProcessPaths, v)
			case "mcp_config_paths":
				a.MCPConfigPaths = append(a.MCPConfigPaths, v)
			case "browser_extension_ids":
				a.BrowserExtensionIDs = append(a.BrowserExtensionIDs, v)
			case "finding_ids":
				a.FindingIDs = append(a.FindingIDs, v)
			}
		}
	}
	return a
}
