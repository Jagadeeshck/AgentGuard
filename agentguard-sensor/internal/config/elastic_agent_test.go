package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestElasticAgentConfigYAMLContainsOutputPath(t *testing.T) {
	y := ElasticAgentConfigYAML()
	if !strings.Contains(y, "output:") || !strings.Contains(y, "path:") {
		t.Fatalf("expected output path in config YAML: %s", y)
	}
}

func TestWriteElasticAgentConfig(t *testing.T) {
	tmp := filepath.Join(t.TempDir(), "config.yml")
	if err := WriteElasticAgentConfig(tmp); err != nil {
		t.Fatalf("WriteElasticAgentConfig failed: %v", err)
	}
	b, err := os.ReadFile(tmp)
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}
	if !strings.Contains(string(b), "collect_prompt_content: false") {
		t.Fatalf("expected privacy guardrails in config: %s", string(b))
	}
}
