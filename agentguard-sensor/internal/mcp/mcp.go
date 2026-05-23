package mcp

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type Server struct {
	Name, Command, ConfigPath string
	Args, Capabilities        []string
}

func candidatePaths() []string {
	home, _ := os.UserHomeDir()
	return []string{filepath.Join(home, ".config/claude/claude_desktop_config.json"), filepath.Join(home, ".cursor/mcp.json"), "mcp.json", ".mcp.json", ".mcp/mcp.json"}
}

func Scan() []Server {
	var out []Server
	for _, p := range candidatePaths() {
		if s, ok := parseFile(p); ok {
			out = append(out, s...)
		}
	}
	return out
}

func parseFile(path string) ([]Server, bool) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, false
	}
	var raw map[string]any
	if json.Unmarshal(b, &raw) != nil {
		return nil, false
	}
	servers := []Server{}
	if m, ok := raw["mcpServers"].(map[string]any); ok {
		for n, v := range m {
			if mv, ok := v.(map[string]any); ok {
				cmd, _ := mv["command"].(string)
				args := toStrings(mv["args"])
				servers = append(servers, Server{Name: n, Command: cmd, Args: args, ConfigPath: path, Capabilities: infer(cmd, args)})
			}
		}
	}
	return servers, true
}
func toStrings(v any) []string {
	a, ok := v.([]any)
	if !ok {
		return nil
	}
	r := []string{}
	for _, x := range a {
		if s, ok := x.(string); ok {
			r = append(r, s)
		}
	}
	return r
}
func infer(cmd string, args []string) []string {
	s := strings.ToLower(cmd + " " + strings.Join(args, " "))
	caps := []string{}
	has := func(k []string) bool {
		for _, w := range k {
			if strings.Contains(s, w) {
				return true
			}
		}
		return false
	}
	if has([]string{"shell", "bash", "zsh", "powershell", "cmd", "terminal", "exec"}) {
		caps = append(caps, "shell")
	}
	if has([]string{"filesystem", "fs", "file", "path", "directory", "root"}) {
		caps = append(caps, "filesystem")
	}
	if has([]string{"browser", "chrome", "playwright", "selenium"}) {
		caps = append(caps, "browser")
	}
	if has([]string{"github", "git"}) {
		caps = append(caps, "code_repository")
	}
	if has([]string{"postgres", "mysql", "sqlite", "mongo", "redis"}) {
		caps = append(caps, "database")
	}
	return caps
}
