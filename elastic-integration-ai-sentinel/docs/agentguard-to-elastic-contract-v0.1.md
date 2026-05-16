# AgentGuard to Elastic Contract v0.1

This contract defines the metadata-only NDJSON event format that a future AgentGuard / AI Sentinel endpoint scanner can write for this Elastic integration. It exists so the Elastic package can be built, linted, tested, and validated before the scanner project exists.

## Boundary

The **future endpoint scanner** is responsible for observing host metadata, assigning `ai_sentinel.finding.type`, calculating risk, redacting sensitive values, and writing one JSON object per line.

This **Elastic integration** is responsible only for collecting those already-produced NDJSON findings with Elastic Agent `filestream`, parsing JSON, normalizing ECS fields, applying defensive redaction for common secret patterns, installing dashboards/rules, and making the data queryable in `logs-ai_sentinel.findings-default`.

This repository must not contain endpoint scanner logic, process enumeration, browser inspection, network capture, prompt collection, clipboard collection, browsing history collection, traffic decryption, or secret storage logic.

## NDJSON format

- One complete JSON object per line.
- UTF-8 encoding.
- No array wrapper.
- No multiline records.
- No comments.
- Each line represents one finding observation or lifecycle update.
- Timestamps use ISO 8601 UTC, for example `2026-05-14T10:00:00Z`.

Minimal shape:

```json
{"@timestamp":"2026-05-14T10:00:00Z","host":{"name":"synthetic-host"},"ai_sentinel":{"finding":{"id":"sample-001","type":"ai_api_connection","name":"AI API connection","status":"open","confidence":0.9},"risk":{"level":"high","score":82,"reasons":["external_ai_api","unknown_process"]},"allowed":false}}
```

## Required ECS fields

Every event should include these ECS fields when known:

- `@timestamp`
- `ecs.version` (optional from producer; the ingest pipeline sets `8.17.0` when absent)
- `event.kind` (optional from producer; pipeline defaults to `alert`)
- `event.category` and `event.type` (optional from producer; pipeline derives known finding types)
- `host.name`
- `process.name` and `process.executable` for process-backed findings
- `user.name` when user attribution is available and non-sensitive
- `file.path` / `file.name` for file-backed findings
- `destination.domain`, `destination.ip`, `destination.address`, and `destination.port` for network/service findings
- `network.transport` and `network.protocol` when known

The pipeline sets these integration fields if absent:

- `event.module: ai_sentinel`
- `event.dataset: ai_sentinel.findings`
- `observer.vendor: AI Sentinel`
- `observer.product: AI Sentinel`
- `observer.type: endpoint`

## Required `ai_sentinel.*` fields

Every finding must include:

- `ai_sentinel.finding.id`: stable finding identifier.
- `ai_sentinel.finding.type`: one supported finding type from this contract.
- `ai_sentinel.finding.name`: human-readable finding name.
- `ai_sentinel.finding.status`: `open`, `acknowledged`, or `closed`.
- `ai_sentinel.finding.confidence`: numeric confidence from `0.0` to `1.0`.
- `ai_sentinel.risk.level`: `low`, `medium`, `high`, or `critical`.
- `ai_sentinel.risk.score`: integer or float from `0` to `100`.
- `ai_sentinel.risk.reasons`: array of normalized reason labels.
- `ai_sentinel.allowed`: boolean allowlist / approved-policy context.

Recommended lifecycle fields:

- `ai_sentinel.finding.first_seen`
- `ai_sentinel.finding.last_seen`
- `ai_sentinel.allowed_by`
- `ai_sentinel.allowed_at`

Type-specific field groups:

- AI/network: `ai_sentinel.ai.provider`, `ai_sentinel.ai.endpoint`, `ai_sentinel.ai.local_service`, `ai_sentinel.ai.model_hint`.
- MCP: `ai_sentinel.mcp.client.name`, `ai_sentinel.mcp.server.name`, `ai_sentinel.mcp.server.command`, `ai_sentinel.mcp.server.args`, `ai_sentinel.mcp.config.path`, `ai_sentinel.mcp.tools`, `ai_sentinel.mcp.capabilities`.
- Browser extension: `ai_sentinel.browser.name`, `ai_sentinel.browser.profile`, `ai_sentinel.extension.id`, `ai_sentinel.extension.name`, `ai_sentinel.extension.version`, `ai_sentinel.extension.permissions`, `ai_sentinel.extension.host_permissions`.
- Startup: `ai_sentinel.startup.type`, `ai_sentinel.startup.name`, `ai_sentinel.startup.path`, `ai_sentinel.startup.command`, `ai_sentinel.startup.enabled`.
- Cyber-agent: `ai_sentinel.cyber_agent.name`, `ai_sentinel.cyber_agent.framework`, `ai_sentinel.cyber_agent.provider`, `ai_sentinel.cyber_agent.activity_type`, `ai_sentinel.cyber_agent.capabilities`, `ai_sentinel.cyber_agent.target_paths`, `ai_sentinel.cyber_agent.security_tools`, `ai_sentinel.cyber_agent.codebase_scan_volume`, `ai_sentinel.cyber_agent.suspicious_keywords`.

## Supported finding types

- `ai_api_connection`
- `mcp_server`
- `mcp_config_modified`
- `browser_extension`
- `startup_item`
- `local_llm_service`
- `suspicious_agent_process`
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

Unknown types should be avoided. If a producer must add a new type, it should first update this contract, fields, tests, detection-rule matrix, and taxonomy documentation.

## Compatibility rules

- Producers may add new optional `ai_sentinel.*` fields without breaking the integration when they are metadata-only and compatible with Elastic mappings.
- Producers must not change the meaning or type of existing fields.
- Risk levels must remain `low`, `medium`, `high`, or `critical`.
- `ai_sentinel.risk.score` must remain numeric and in the `0-100` range.
- Arrays should stay arrays even when only one value is present, for example `ai_sentinel.risk.reasons` and `ai_sentinel.mcp.capabilities`.
- Breaking schema changes require a new contract document version.

## Privacy requirements

The producer must emit only metadata needed for detection, triage, dashboards, and rules. It must not emit:

- Raw prompt text, completion text, conversation transcripts, or private tool-call payloads.
- Clipboard content.
- Browsing history, page contents, form contents, or email contents.
- Decrypted network traffic.
- API keys, OAuth tokens, passwords, cookies, session IDs, private keys, or other secrets.
- Private source file contents or exploit payload contents.

Allowed metadata includes process names, executable paths, destination domains, destination ports, tool names, browser extension IDs, MCP server names, normalized capabilities, aggregate scan counts, and non-sensitive model/provider hints.

## Redaction requirements

Producers must redact secrets before writing NDJSON. The ingest pipeline also redacts common patterns defensively, but pipeline redaction is not a substitute for producer-side minimization.

Use `[REDACTED]` for sensitive values while preserving detection context:

- Allowed: `/opt/agent --provider openai --api-key [REDACTED]`
- Allowed: `https://api.example.invalid/v1/responses?api_key=[REDACTED]`
- Not allowed: raw API keys, bearer tokens, passwords, cookies, session identifiers, or private key material.

## Finding lifecycle

Use `ai_sentinel.finding.status` consistently:

- `open`: current finding requiring attention.
- `acknowledged`: reviewed but still present.
- `closed`: no longer present or accepted through policy.

Use `ai_sentinel.allowed: true` only when the finding matches an approved allowlist or policy exception. Include `ai_sentinel.allowed_by` and `ai_sentinel.allowed_at` when available.
