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
| `ai_sentinel.finding.type` | keyword | Yes | Finding family. Supported MVP values are `ai_api_connection`, `mcp_server`, `browser_extension`, `startup_item`, and `local_llm_service`. |
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
