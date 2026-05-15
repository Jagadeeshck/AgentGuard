# Detection Rule Test Matrix

This matrix maps packaged Elastic Security rules to finding types, sample events, and expected validation outcomes. It is a documentation-level test plan for independent package validation before an endpoint scanner exists.

| Rule file | Primary finding types | Positive sample IDs | Expected behavior | Tuning dimensions |
| --- | --- | --- | --- | --- |
| `unknown_process_ai_api.toml` | `ai_api_connection` | `sample-ai-api-unknown-process` | Alerts when an unapproved process contacts an AI API endpoint. | Process executable, provider, host group. |
| `multiple_ai_api_connections_threshold.toml` | `ai_api_connection` | `sample-ai-api-unknown-process` | Alerts on repeated AI API connections above threshold. | Provider, host group, approved automation. |
| `untrusted_mcp_shell_filesystem.toml` | `mcp_server` | `sample-mcp-shell-filesystem` | Alerts on untrusted MCP with shell and filesystem capabilities. | MCP server name, command path, restricted root. |
| `new_modified_mcp_config.toml` | `mcp_config_modified` | `sample-mcp-config-modified` | Alerts on new or modified MCP configuration metadata. | Config path, managed deployment policy. |
| `ai_browser_extension_broad_permissions.toml` | `browser_extension` | `sample-browser-extension-broad-permissions` | Alerts on broad extension permissions or host permissions. | Extension ID, version, managed profile. |
| `ai_tool_added_startup.toml` | `startup_item` | `sample-startup-item-ai-helper` | Alerts when AI-related helper is configured for startup. | Startup item name, path, software owner. |
| `local_llm_exposed.toml` | `local_llm_service` | `sample-local-llm-exposed` | Alerts on exposed local LLM listener metadata. | Loopback-only services, approved host group. |
| `ai_sentinel_critical_finding.toml` | All finding types | `sample-cyber-agent-general`, `sample-exploit-development` | Alerts on critical risk findings. | Validated allowlist status and risk reason quality. |
| `critical_ai_cyber_agent_activity.toml` | Cyber-agent finding types | `sample-cyber-agent-general`, `sample-exploit-development` | Alerts on critical cyber-agent behavior combinations. | Authorized security research workstations. |
| `untrusted_ai_agent_mcp_browser_shell.toml` | `ai_cyber_agent_activity`, `ai_agent_shell_tool_use` | `sample-cyber-agent-general`, `sample-agent-shell-tool-use` | Alerts on untrusted agent combining MCP/browser/shell-style capabilities. | Agent framework, approved capabilities. |
| `untrusted_ai_agent_security_tools.toml` | `ai_security_tool_mcp_server` | `sample-security-tool-mcp` | Alerts on untrusted AI agent access to security tools through MCP. | Approved tools, team ownership, ticket metadata. |
| `ai_agent_shell_filesystem_mcp.toml` | `ai_agent_shell_tool_use`, `ai_security_tool_mcp_server` | `sample-agent-shell-tool-use`, `sample-security-tool-mcp` | Alerts on combined shell, filesystem, and MCP capabilities. | Project root restriction, command path. |
| `ai_agent_sensitive_source_scan.toml` | `ai_agent_sensitive_repo_scan` | `sample-sensitive-repo-scan` | Alerts on scanning sensitive repository areas. | Approved repo paths, user group. |
| `ai_agent_sandbox_escape_research.toml` | `ai_sandbox_escape_research` | `sample-sandbox-escape-research` | Alerts on sandbox escape research indicators. | Authorized container security research. |
| `ai_agent_exploit_like_files.toml` | `ai_exploit_development_activity` | `sample-exploit-development` | Alerts on exploit-like file/project metadata. | Authorized exploit research projects. |
| `possible_mythos_like_vulnerability_research_agent.toml` | `ai_vulnerability_research_agent`, `ai_cyber_agent_activity` | `sample-vulnerability-research-agent` | Alerts only when mythos-like labels appear with concrete risky behaviors. | Do not alert on name alone; require behavior. |

## Manual validation checklist

- Confirm every rule maps to at least one synthetic positive sample.
- Confirm each sample uses metadata-only fields.
- Confirm allowlisted safe scenarios are documented separately rather than encoded as real secrets or private content.
- Confirm pipeline tests pass before asset tests are run.
