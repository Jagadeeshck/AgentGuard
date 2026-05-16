# Detection Rule Test Matrix

This matrix maps the synthetic NDJSON validation corpus in `test-data/` to the expected Elastic Security rules. It is a manual validation aid for `elastic-package test asset` and local lab testing; sample data is synthetic and contains no real secrets, prompt content, browsing history, clipboard content, or private file contents.

| Sample file | Finding type | Expected rule | Expected severity | Should alert yes/no | Notes |
|---|---|---|---|---|---|
| `test-data/ai-api-connection-high.ndjson` | `ai_api_connection` | `unknown_process_ai_api.toml` | high | yes | Unknown process path connects to an external AI API endpoint. |
| `test-data/ai-api-connection-high.ndjson` | `ai_api_connection` | `multiple_ai_api_connections_threshold.toml` | medium | no | Single sample line alone should not meet the threshold rule. |
| `test-data/mcp-server-dangerous.ndjson` | `mcp_server` | `untrusted_mcp_shell_filesystem.toml` | high | yes | Untrusted MCP server exposes shell and filesystem capabilities. |
| `test-data/mcp-server-dangerous.ndjson` | `mcp_server` | `ai_agent_shell_filesystem_mcp.toml` | high | no | Rule is scoped to cyber-agent findings; MCP server risk is covered by `untrusted_mcp_shell_filesystem.toml`. |
| `test-data/cyber-agent-activity.ndjson` | `ai_cyber_agent_activity` | `critical_ai_cyber_agent_activity.toml` | critical | yes | Untrusted agent combines MCP, shell, filesystem, browser, and security tooling. |
| `test-data/browser-extension-risk.ndjson` | `browser_extension` | `ai_browser_extension_broad_permissions.toml` | high | yes | Extension has `all_urls` and `nativeMessaging`. |
| `test-data/local-llm-exposed.ndjson` | `local_llm_service` | `local_llm_exposed.toml` | high | yes | Local LLM service is bound to `0.0.0.0`. |
| `test-data/safe-ollama-loopback.ndjson` | `local_llm_service` | `local_llm_exposed.toml` | low | no | Loopback-only and allowlisted. |
| `test-data/safe-approved-ai-client.ndjson` | `ai_api_connection` | `unknown_process_ai_api.toml` | low | no | Approved internal AI API client with allowlist metadata. |
| `test-data/startup-item-ai-agent.ndjson` | `startup_item` | `ai_tool_added_startup.toml` | high | yes | AI helper creates startup persistence. |
| `test-data/startup-item-ai-agent.ndjson` | `startup_item` | `ai_sentinel_critical_finding.toml` | critical | no | High severity only; should not match critical catch-all. |
| `test-data/sensitive-repo-scan.ndjson` | `ai_agent_sensitive_repo_scan` | `ai_agent_sensitive_source_scan.toml` | high | yes | Agent scans auth, token, and CI/CD metadata paths. |
| `test-data/sensitive-repo-scan.ndjson` | `ai_agent_mass_codebase_analysis` | `untrusted_ai_agent_mcp_browser_shell.toml` | high | no | This sample is sensitive repo scanning, not shell/browser/MCP activity. |
| `test-data/exploit-like-file-write.ndjson` | `ai_exploit_development_activity` | `ai_agent_exploit_like_files.toml` | critical | yes | Untrusted agent writes exploit-like synthetic PoC files. |
| `test-data/exploit-like-file-write.ndjson` | `ai_exploit_development_activity` | `ai_sentinel_critical_finding.toml` | critical | yes | Score is in the critical range. |

## Validation checklist

- Every expected alert maps to a synthetic positive sample.
- Every safe sample documents why it should not alert.
- Rules should remain behavior-based and should not depend on prompt content, secrets, clipboard content, browsing history, decrypted traffic, or private file contents.
- Pipeline tests must pass before asset tests are run.
