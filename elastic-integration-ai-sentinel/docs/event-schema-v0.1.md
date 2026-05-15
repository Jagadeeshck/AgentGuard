# AI Sentinel finding event schema v0.1

This document defines the minimum event contract consumed by the AI Sentinel Elastic integration package. The package ingests newline-delimited JSON findings emitted by a separate AI Sentinel endpoint product; it does not implement endpoint scanning.

## Scope and privacy boundaries

The integration is designed for security metadata only. Producers **must not** emit prompt content, model responses, secrets, decrypted traffic, clipboard content, or browsing history. The ingest pipeline defensively drops several known prohibited payload fields when they appear accidentally, but upstream producers remain responsible for avoiding collection.

All non-ECS custom fields must live below `ai_sentinel.*`. ECS fields such as `host.*`, `process.*`, `user.*`, `file.*`, `source.*`, `destination.*`, `network.*`, `dns.*`, `event.*`, `observer.*`, and `related.*` may be used when they match ECS semantics.

## Data stream

- Type: `logs`
- Dataset: `ai_sentinel.findings`
- Default data stream: `logs-ai_sentinel.findings-default`

Each line must be a single JSON object. The package reads each line with Elastic Agent `filestream` and parses it in the `findings` ingest pipeline.

## Required fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `@timestamp` | date | Recommended | Time the finding was observed. Elastic will provide an ingest timestamp if omitted. |
| `ai_sentinel.finding.id` | keyword | Yes | Stable finding identifier. |
| `ai_sentinel.finding.type` | keyword | Yes | Finding family. Supported values include `ai_api_connection`, `mcp_server`, `browser_extension`, `startup_item`, `local_llm_service`, `suspicious_agent_process`, `mcp_config_modified`, and the AI cyber-agent detection pack values: `ai_cyber_agent_activity`, `ai_vulnerability_research_agent`, `ai_sandbox_escape_research`, `ai_fuzzing_activity`, `ai_reverse_engineering_activity`, `ai_exploit_development_activity`, `ai_security_tool_mcp_server`, `ai_agent_shell_tool_use`, `ai_agent_sensitive_repo_scan`, and `ai_agent_mass_codebase_analysis`. |
| `ai_sentinel.risk.level` | keyword | Yes | One of `low`, `medium`, `high`, or `critical`. |
| `ai_sentinel.risk.score` | float | Yes | Numeric risk score, typically 0-100. |
| `ai_sentinel.allowed` | boolean | Recommended | Whether the finding is allowlisted or approved by policy. |

## Common ECS fields

Use ECS fields for host, user, process, file, and network context when available:

- `host.name`, `host.id`, `host.os.type`
- `user.name`
- `process.name`, `process.executable`, `process.command_line`, `process.pid`, `process.parent.name`
- `file.path`, `file.name`
- `destination.ip`, `destination.address`, `destination.domain`, `destination.port`
- `source.ip`, `source.port`
- `network.transport`, `network.protocol`
- `dns.question.name`

The pipeline adds `event.module`, `event.dataset`, `observer.*`, AI Sentinel tags, `related.*`, and derived event categorization fields.

## Finding-specific fields

### `ai_api_connection`

Represents metadata that a process connected to a known or suspected AI API endpoint.

Recommended fields:

- `ai_sentinel.ai.provider`
- `ai_sentinel.ai.endpoint` with query-string secrets removed or redacted
- `ai_sentinel.ai.local_service: false`
- `ai_sentinel.ai.confidence`
- ECS `process.*`, `destination.*`, and `network.*`

Pipeline-derived fields:

- `event.category: ["network"]`
- `event.type: ["connection"]`

### `mcp_server`

Represents metadata about a configured or discovered MCP server.

Recommended fields:

- `ai_sentinel.mcp.client.name`
- `ai_sentinel.mcp.server.name`
- `ai_sentinel.mcp.server.command` with secrets redacted
- `ai_sentinel.mcp.server.args` with secrets redacted
- `ai_sentinel.mcp.config.path`
- `ai_sentinel.mcp.tools`
- `ai_sentinel.mcp.capabilities`
- `ai_sentinel.mcp.risk_reasons`

Pipeline-derived fields:

- `event.category: ["configuration"]`
- `event.type: ["info"]`

### `browser_extension`

Represents metadata about an AI-related browser extension or an extension with risky AI-adjacent capabilities.

Recommended fields:

- `ai_sentinel.browser.name`
- `ai_sentinel.browser.profile`
- `ai_sentinel.extension.id`
- `ai_sentinel.extension.name`
- `ai_sentinel.extension.version`
- `ai_sentinel.extension.permissions`
- `ai_sentinel.extension.host_permissions`
- `ai_sentinel.extension.risk_reasons`

Do not emit browsing history, page content, form values, cookies, or clipboard content.

Pipeline-derived fields:

- `event.category: ["configuration"]`
- `event.type: ["info"]`

### `startup_item`

Represents metadata about an AI tool, MCP server, or related helper configured to run at startup.

Recommended fields:

- `ai_sentinel.startup.type`
- `ai_sentinel.startup.name`
- `ai_sentinel.startup.path`
- `ai_sentinel.startup.command` with secrets redacted
- `ai_sentinel.startup.enabled`
- ECS `file.*` when applicable

Pipeline-derived fields:

- `event.category: ["configuration"]`
- `event.type: ["creation"]`

### `local_llm_service`

Represents metadata about a local LLM service listener.

Recommended fields:

- `ai_sentinel.ai.provider`
- `ai_sentinel.ai.endpoint`
- `ai_sentinel.ai.local_service: true`
- `ai_sentinel.ai.model_hint` if it does not reveal prompt or response content
- ECS `process.*`, `destination.*`, and `network.*`

Pipeline-derived fields:

- `event.category: ["network"]`
- `event.type: ["info"]`

## AI cyber-agent detection pack

The cyber-agent pack detects Mythos-like AI cyber-agent activity using behaviour-based metadata. The integration must not collect prompt content, secrets, decrypted traffic, clipboard content, browsing history, exploit logic, or offensive payloads. The optional word `mythos` may appear as one weak indicator in `ai_sentinel.cyber_agent.suspicious_keywords`, but producers and rules must combine it with concrete behaviours such as shell access, filesystem access, security tooling, sensitive path scanning, or exploit-development file metadata.

Supported cyber-agent finding types:

- `ai_cyber_agent_activity`
- `ai_vulnerability_research_agent`
- `ai_sandbox_escape_research`
- `ai_fuzzing_activity`
- `ai_reverse_engineering_activity`
- `ai_exploit_development_activity`
- `ai_security_tool_mcp_server`
- `ai_agent_shell_tool_use`
- `ai_agent_sensitive_repo_scan`
- `ai_agent_mass_codebase_analysis`

Recommended fields under `ai_sentinel.cyber_agent.*`:

- `name`: observed or declared agent name, when available.
- `framework`: non-sensitive framework/runtime hint, such as an agent framework name.
- `model_hint`: non-sensitive model family hint.
- `provider`: AI provider associated with the agent.
- `activity_type`: behaviour classification such as `security_tool_use`, `vulnerability_research`, `sandbox_escape_research`, `fuzzing`, `reverse_engineering`, `exploit_development`, `sensitive_repo_scan`, or `mass_codebase_analysis`.
- `capabilities`: observed capabilities such as `mcp`, `shell`, `filesystem`, `browser`, `fuzzing`, or `reverse_engineering`.
- `target_paths`: paths or repository areas inspected or written by the agent; include paths only, not contents.
- `security_tools`: names of security tools invoked or configured by the agent.
- `codebase_scan_volume`: approximate count of files, symbols, or repository objects inspected.
- `suspicious_keywords`: behavioural metadata keywords from filenames, commands, or normalized findings; do not store prompt content.
- `mitre_tactics` and `mitre_techniques`: ATT&CK tactic and technique identifiers associated with the finding.

Pipeline-derived fields:

- Cyber-agent findings receive the `cyber_agent_activity` tag.
- Most cyber-agent findings receive `event.category: ["process"]` and `event.type: ["info"]`.
- `ai_agent_sensitive_repo_scan` receives `event.category: ["file"]`.
- `ai_exploit_development_activity` receives `event.type: ["creation"]`.

## Risk mapping

The ingest pipeline copies `ai_sentinel.risk.score` to `event.risk_score` and derives `event.severity` from `ai_sentinel.risk.level`:

| Risk level | Event severity |
| --- | ---: |
| `low` | 21 |
| `medium` | 47 |
| `high` | 73 |
| `critical` | 99 |

## Example event

```json
{
  "@timestamp": "2026-05-14T10:00:00Z",
  "host": { "name": "host-a" },
  "process": {
    "name": "python",
    "executable": "/usr/bin/python3",
    "command_line": "python app.py --api-key [REDACTED]"
  },
  "destination": {
    "domain": "api.openai.com",
    "address": "api.openai.com",
    "port": 443
  },
  "network": { "transport": "tcp", "protocol": "tls" },
  "ai_sentinel": {
    "finding": {
      "id": "f-ai-api-1",
      "type": "ai_api_connection",
      "name": "AI API connection",
      "status": "open",
      "confidence": 0.9
    },
    "risk": {
      "level": "high",
      "score": 82,
      "reasons": ["unknown_process"]
    },
    "ai": {
      "provider": "openai",
      "endpoint": "https://api.openai.com/v1/responses",
      "local_service": false,
      "confidence": 0.95
    },
    "allowed": false
  }
}
```
