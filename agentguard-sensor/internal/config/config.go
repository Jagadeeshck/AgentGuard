package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type Config struct {
	Output    OutputConfig
	Watch     WatchConfig
	Scan      ScanConfig
	Privacy   PrivacyConfig
	Allowlist AllowlistConfig
	Logging   LoggingConfig
}
type OutputConfig struct{ Path, Format string }
type WatchConfig struct {
	Interval     time.Duration
	StateFile    string
	EmitResolved bool
}
type ScanConfig struct{ MCP, BrowserExtensions, LocalLLM, StartupItems, Processes bool }
type PrivacyConfig struct{ RedactSecrets, CollectPromptContent, CollectClipboard, CollectBrowserHistory, CollectPrivateFileContents, DecryptTraffic bool }
type AllowlistConfig struct{ Path string }
type LoggingConfig struct{ Level, Path string }

func DefaultConfig() Config {
	c := Config{Output: OutputConfig{Format: "ndjson"}, Watch: WatchConfig{Interval: 60 * time.Second, EmitResolved: true}, Scan: ScanConfig{true, true, true, true, true}, Privacy: PrivacyConfig{RedactSecrets: true}, Logging: LoggingConfig{Level: "info"}}
	switch runtime.GOOS {
	case "darwin":
		c.Output.Path = "/Library/Logs/AgentGuard/findings.ndjson"
		c.Watch.StateFile = "/var/lib/agentguard/state.json"
		c.Allowlist.Path = "/Library/Application Support/AgentGuard/allowlist.yml"
		c.Logging.Path = "/Library/Logs/AgentGuard/agentguard.log"
	case "windows":
		c.Output.Path = `C:\ProgramData\AgentGuard\logs\findings.ndjson`
		c.Watch.StateFile = `C:\ProgramData\AgentGuard\state.json`
		c.Allowlist.Path = `C:\ProgramData\AgentGuard\allowlist.yml`
		c.Logging.Path = `C:\ProgramData\AgentGuard\agentguard.log`
	default:
		c.Output.Path = "/var/log/agentguard/findings.ndjson"
		c.Watch.StateFile = "/var/lib/agentguard/state.json"
		c.Allowlist.Path = "/etc/agentguard/allowlist.yml"
		c.Logging.Path = "/var/log/agentguard/agentguard.log"
	}
	return c
}
func DefaultConfigPath() string {
	if runtime.GOOS == "darwin" {
		return "/Library/Application Support/AgentGuard/config.yml"
	}
	if runtime.GOOS == "windows" {
		return `C:\ProgramData\AgentGuard\config.yml`
	}
	return "/etc/agentguard/config.yml"
}

func Load(path string) (Config, error) {
	c := DefaultConfig()
	if path == "" {
		return c, nil
	}
	f, err := os.Open(path)
	if err != nil {
		return c, err
	}
	defer f.Close()
	sec := ""
	s := bufio.NewScanner(f)
	for s.Scan() {
		ln := strings.TrimSpace(s.Text())
		if ln == "" || strings.HasPrefix(ln, "#") {
			continue
		}
		if strings.HasSuffix(ln, ":") && !strings.Contains(ln, " ") {
			sec = strings.TrimSuffix(ln, ":")
			continue
		}
		kv := strings.SplitN(ln, ":", 2)
		if len(kv) != 2 {
			continue
		}
		k := strings.TrimSpace(kv[0])
		v := strings.Trim(strings.TrimSpace(kv[1]), "\"")
		apply(&c, sec, k, v)
	}
	if err := s.Err(); err != nil {
		return c, err
	}
	return c, c.Validate()
}

func apply(c *Config, sec, k, v string) {
	b := strings.EqualFold(v, "true")
	switch sec {
	case "output":
		if k == "path" {
			c.Output.Path = v
		}
		if k == "format" {
			c.Output.Format = v
		}
	case "watch":
		if k == "state_file" {
			c.Watch.StateFile = v
		}
		if k == "emit_resolved" {
			c.Watch.EmitResolved = b
		}
		if k == "interval" {
			if d, err := time.ParseDuration(v); err == nil {
				c.Watch.Interval = d
			}
		}
	case "scan":
		if k == "mcp" {
			c.Scan.MCP = b
		}
		if k == "browser_extensions" {
			c.Scan.BrowserExtensions = b
		}
		if k == "local_llm" {
			c.Scan.LocalLLM = b
		}
		if k == "startup_items" {
			c.Scan.StartupItems = b
		}
		if k == "processes" {
			c.Scan.Processes = b
		}
	case "privacy":
		if k == "redact_secrets" {
			c.Privacy.RedactSecrets = b
		}
		if k == "collect_prompt_content" {
			c.Privacy.CollectPromptContent = b
		}
		if k == "collect_clipboard" {
			c.Privacy.CollectClipboard = b
		}
		if k == "collect_browser_history" {
			c.Privacy.CollectBrowserHistory = b
		}
		if k == "collect_private_file_contents" {
			c.Privacy.CollectPrivateFileContents = b
		}
		if k == "decrypt_traffic" {
			c.Privacy.DecryptTraffic = b
		}
	case "allowlist":
		if k == "path" {
			c.Allowlist.Path = v
		}
	case "logging":
		if k == "level" {
			c.Logging.Level = v
		}
		if k == "path" {
			c.Logging.Path = v
		}
	}
}
func (c Config) Validate() error {
	if c.Output.Format != "" && c.Output.Format != "ndjson" {
		return errors.New("output.format must be ndjson")
	}
	if c.Privacy.CollectPromptContent || c.Privacy.CollectClipboard || c.Privacy.CollectBrowserHistory || c.Privacy.CollectPrivateFileContents || c.Privacy.DecryptTraffic {
		return errors.New("privacy guardrails are enforced in MVP and cannot be enabled")
	}
	return nil
}
func EnsureRuntimePaths(c Config) error {
	for _, p := range []string{filepath.Dir(c.Output.Path), filepath.Dir(c.Watch.StateFile)} {
		if p == "" || p == "." {
			continue
		}
		if err := os.MkdirAll(p, 0o755); err != nil {
			return err
		}
	}
	f, err := os.OpenFile(c.Output.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("output path not writable: %w", err)
	}
	_ = f.Close()
	return nil
}
