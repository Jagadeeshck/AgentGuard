package browser

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type Extension struct {
	ID, Name, Description, Path string
	Risky                       bool
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
	for _, k := range []string{"<all_urls>", "tabs", "activetab", "scripting", "webrequest", "clipboardread", "nativemessaging"} {
		if strings.Contains(s, k) {
			return true
		}
	}
	return false
}
func flat(v any) string { b, _ := json.Marshal(v); return string(b) }
