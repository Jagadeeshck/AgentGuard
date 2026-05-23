package browser

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type Extension struct {
	ID, Name, Description, Path, Browser, Profile string
	Risky                                         bool
}

func Scan() []Extension {
	home, _ := os.UserHomeDir()
	paths := []struct{ browser, profile, root string }{
		{"chrome", "default", filepath.Join(home, ".config/google-chrome/Default/Extensions")},
		{"chromium", "default", filepath.Join(home, ".config/chromium/Default/Extensions")},
	}
	out := []Extension{}
	for _, p := range paths {
		_ = filepath.Walk(p.root, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() || filepath.Base(path) != "manifest.json" {
				return nil
			}
			ext, ok := ParseManifest(path)
			if !ok {
				return nil
			}
			ext.Browser = p.browser
			ext.Profile = p.profile
			if ext.Risky {
				out = append(out, ext)
			}
			return nil
		})
	}
	return out
}

func ParseManifest(path string) (Extension, bool) {
	b, err := os.ReadFile(path)
	if err != nil {
		return Extension{}, false
	}
	var m map[string]any
	if json.Unmarshal(b, &m) != nil {
		return Extension{}, false
	}
	name, _ := m["name"].(string)
	desc, _ := m["description"].(string)
	id := filepath.Base(filepath.Dir(path))
	risky := hasRisk(m)
	return Extension{ID: id, Name: name, Description: desc, Path: path, Risky: risky}, true
}
func hasRisk(m map[string]any) bool {
	s := strings.ToLower(flat(m))
	for _, k := range []string{"<all_urls>", "tabs", "activetab", "scripting", "webrequest", "nativemessaging"} {
		if strings.Contains(s, k) {
			return true
		}
	}
	return false
}
func flat(v any) string { b, _ := json.Marshal(v); return string(b) }
