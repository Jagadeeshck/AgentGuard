package startup

import (
	"os"
	"path/filepath"
	"strings"
)

type Item struct {
	Path  string
	Match bool
}

var keywords = []string{"claude", "cursor", "ollama", "lmstudio", "mcp", "agent", "langchain", "llamaindex", "autogpt", "crewai", "open-interpreter", "ai-helper", "copilot"}

func Scan() []Item {
	home, _ := os.UserHomeDir()
	paths := []string{filepath.Join(home, ".config/autostart"), "/etc/xdg/autostart", filepath.Join(home, ".config/systemd/user"), "/etc/systemd/system"}
	out := []Item{}
	for _, p := range paths {
		_ = filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}
			l := strings.ToLower(path)
			m := false
			for _, k := range keywords {
				if strings.Contains(l, k) {
					m = true
					break
				}
			}
			if m {
				out = append(out, Item{Path: path, Match: true})
			}
			return nil
		})
	}
	return out
}
