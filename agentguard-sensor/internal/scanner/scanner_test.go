package scanner

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/agentguard/agentguard-sensor/internal/browser"
	"github.com/agentguard/agentguard-sensor/internal/localai"
	"github.com/agentguard/agentguard-sensor/internal/mcp"
	"github.com/agentguard/agentguard-sensor/internal/process"
	"github.com/agentguard/agentguard-sensor/internal/startup"
)

func withFakeScanners(t *testing.T) {
	t.Helper()
	oldMCP, oldProc, oldLocal, oldStartup, oldBrowser := mcpScan, processScan, localaiScan, startupScan, browserScan
	mcpScan = func() []mcp.Server {
		return []mcp.Server{{Name: "srv", Command: "bash", Args: []string{"--root", "/tmp"}, ConfigPath: "/mcp.json", Capabilities: []string{"shell", "filesystem"}}}
	}
	processScan = func() []process.Proc {
		return []process.Proc{{Name: "node", PID: 1, PPID: 0, Executable: "/bin/node", CommandLine: "node cursor --token=secret"}}
	}
	localaiScan = func() []localai.Service {
		return []localai.Service{{Addr: "0.0.0.0", Port: 11434, ProcessName: "ollama", Exposed: true}}
	}
	startupScan = func() []startup.Item {
		return []startup.Item{{Path: "/startup/agent.desktop", Command: "agent", Match: true}}
	}
	browserScan = func() []browser.Extension {
		return []browser.Extension{{Browser: "chrome", Profile: "Default", ID: "ext1", Name: "AI Helper", Path: "/ext"}}
	}
	t.Cleanup(func() {
		mcpScan, processScan, localaiScan, startupScan, browserScan = oldMCP, oldProc, oldLocal, oldStartup, oldBrowser
	})
}

func TestScanFindingsAllowlistBranches(t *testing.T) {
	withFakeScanners(t)
	path := filepath.Join(t.TempDir(), "allow.yml")
	content := "mcp_config_paths: [/mcp.json]\nprocess_paths: [/bin/node]\nlocal_ports: [11434]\nbrowser_extension_ids: [ext1]\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	events, _ := ScanFindings(path)
	allowed := map[string]bool{}
	for _, e := range events {
		allowed[e.AgentGuard.Finding.Type] = e.AgentGuard.Allowed
	}
	for _, kind := range []string{"mcp_server", "suspicious_agent_process", "local_llm_service", "browser_extension"} {
		if !allowed[kind] {
			t.Fatalf("%s was not allowlisted: %#v", kind, allowed)
		}
	}
	if allowed["startup_item"] {
		t.Fatal("startup_item should not be allowlisted")
	}
}

func TestScanFindingsNotAllowed(t *testing.T) {
	withFakeScanners(t)
	events, _ := ScanFindings("")
	if len(events) != 5 {
		t.Fatalf("events=%d", len(events))
	}
	for _, e := range events {
		if e.AgentGuard.Allowed {
			t.Fatalf("unexpected allowed event: %#v", e)
		}
	}
}

func TestLoadAllowlistEdgeCases(t *testing.T) {
	dir := t.TempDir()
	empty := filepath.Join(dir, "empty.yml")
	if err := os.WriteFile(empty, nil, 0o644); err != nil {
		t.Fatal(err)
	}
	if got := loadAllowlist(empty); len(got.ProcessPaths) != 0 || len(got.LocalPorts) != 0 {
		t.Fatalf("empty got %#v", got)
	}
	malformed := filepath.Join(dir, "bad.yml")
	if err := os.WriteFile(malformed, []byte("local_ports: [abc, 8080]\nignored line\n- orphan\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	got := loadAllowlist(malformed)
	if len(got.LocalPorts) != 1 || got.LocalPorts[0] != 8080 {
		t.Fatalf("malformed got %#v", got)
	}
}

func TestListKinds(t *testing.T) {
	withFakeScanners(t)
	for _, kind := range []string{"mcp", "extensions", "local-ai", "startup", "processes"} {
		got, err := List(kind)
		if err != nil || len(got) == 0 {
			t.Fatalf("%s got %#v err %v", kind, got, err)
		}
	}
	if _, err := List("bogus"); err == nil {
		t.Fatal("expected error")
	}
}
