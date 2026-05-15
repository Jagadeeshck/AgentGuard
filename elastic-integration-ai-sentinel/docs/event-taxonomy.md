# Event Taxonomy

The `ai_sentinel.finding.type` field is the primary taxonomy key. The ingest pipeline derives ECS categorization, tags, risk fields, and observability metadata from these values.

| Finding type | Purpose | Primary ECS category | Primary ECS type | Typical risk drivers |
| --- | --- | --- | --- | --- |
| `ai_api_connection` | Process connected to a known or suspected AI API endpoint. | `network` | `connection` | Unknown process, unusual destination, unapproved provider. |
| `mcp_server` | MCP server configured or discovered. | `configuration` | `info` | Shell capability, filesystem capability, untrusted command path. |
| `browser_extension` | AI-related or risky browser extension metadata. | `configuration` | `info` | Broad host permissions, native messaging, unapproved extension. |
| `startup_item` | AI tool, helper, or MCP bridge configured for startup. | `configuration` | `creation` | Persistence, untrusted path, unexpected user scope. |
| `local_llm_service` | Local LLM service listener metadata. | `network` | `info` | Non-loopback bind, unexpected port, unapproved service owner. |
| `mcp_config_modified` | MCP configuration file changed. | `configuration` | `info` | New server, expanded tools, sensitive config path. |
| `suspicious_agent_process` | Agent-like process behavior requiring review. | `process` | `info` | Temporary path, unusual command, combined AI and automation indicators. |
| `ai_cyber_agent_activity` | General behavior-based AI cyber-agent activity. | `process` | `info` | Multiple risky capabilities combined with untrusted agent context. |
| `ai_vulnerability_research_agent` | AI-assisted vulnerability research metadata. | `process` | `info` | Vulnerability keywords plus repo analysis and security tooling. |
| `ai_sandbox_escape_research` | Sandbox/container escape research indicators. | `process` | `info` | Container-sensitive paths, sandbox terms, runtime introspection. |
| `ai_fuzzing_activity` | AI-assisted fuzzing activity metadata. | `process` | `info` | Fuzzer tools, corpus paths, crash triage metadata. |
| `ai_reverse_engineering_activity` | AI-assisted reverse engineering metadata. | `process` | `info` | Disassembler tools, binary target paths, reverse engineering labels. |
| `ai_exploit_development_activity` | Exploit-like file or project metadata. | `process` | `creation` | PoC filenames, exploit-development labels, untrusted agent writes. |
| `ai_security_tool_mcp_server` | MCP server exposes security tools to an AI agent. | `process` | `info` | MCP plus security tools plus untrusted agent context. |
| `ai_agent_shell_tool_use` | AI agent has or uses shell execution capability. | `process` | `info` | Shell capability combined with filesystem, browser, or MCP access. |
| `ai_agent_sensitive_repo_scan` | AI agent scans sensitive repository areas. | `file` | `info` | Auth, credential, CI/CD, infrastructure, or secret-adjacent paths. |
| `ai_agent_mass_codebase_analysis` | AI agent performs large codebase analysis. | `process` | `info` | High scan volume, broad target paths, unknown business justification. |

## Taxonomy principles

- Prefer behavior labels over prompt-derived labels.
- Use `ai_sentinel.cyber_agent.suspicious_keywords` only for normalized metadata keywords from filenames, commands, paths, or rule observations.
- Do not place raw prompts, completions, code contents, file contents, or secrets into taxonomy fields.
- Combine weak indicators with concrete behaviors before assigning high or critical risk.
