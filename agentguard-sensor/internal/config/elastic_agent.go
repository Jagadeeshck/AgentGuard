package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func ElasticAgentConfigYAML() string {
	c := DefaultConfig()
	return fmt.Sprintf(`output:
  path: %s
  format: ndjson
watch:
  interval: %s
  state_file: %s
  emit_resolved: %t
scan:
  mcp: %t
  browser_extensions: %t
  local_llm: %t
  startup_items: %t
  processes: %t
privacy:
  redact_secrets: %t
  collect_prompt_content: false
  collect_clipboard: false
  collect_browser_history: false
  collect_private_file_contents: false
  decrypt_traffic: false
allowlist:
  path: %s
logging:
  level: %s
  path: %s
`, c.Output.Path, c.Watch.Interval, c.Watch.StateFile, c.Watch.EmitResolved, c.Scan.MCP, c.Scan.BrowserExtensions, c.Scan.LocalLLM, c.Scan.StartupItems, c.Scan.Processes, c.Privacy.RedactSecrets, c.Allowlist.Path, c.Logging.Level, c.Logging.Path)
}

func WriteElasticAgentConfig(outputPath string) error {
	if outputPath == "" {
		outputPath = DefaultConfigPath()
	}
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return err
	}
	return os.WriteFile(outputPath, []byte(ElasticAgentConfigYAML()), 0o644)
}
