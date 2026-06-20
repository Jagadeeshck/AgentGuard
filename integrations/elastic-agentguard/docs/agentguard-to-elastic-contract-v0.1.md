# AgentGuard to Elastic Contract v0.1

This contract defines the metadata-only NDJSON event format that a future AgentGuard endpoint scanner can write for this Elastic integration. It exists so the Elastic package can be built, linted, tested, and validated before the scanner project exists.

## Boundary

The **future endpoint scanner** is responsible for observing host metadata, assigning `agentguard.finding.type`, calculating risk, redacting sensitive values, and writing one JSON object per line.

This **Elastic integration** is responsible only for collecting those already-produced NDJSON findings with Elastic Agent `filestream`, parsing JSON, normalizing ECS fields, applying defensive redaction for common secret patterns, keeping development dashboard/rule references, and making the data queryable in `logs-agentguard.findings-default`.

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
{"@timestamp":"2026-05-14T10:00:00Z","host":{"name":"synthetic-host"},"agentguard":{"finding":{"id":"sample-001","type":"ai_api_connection","name":"AI API connection","status":"open","confidence":0.9},"risk":{"level":"high","score":82,"reasons":["external_ai_api","unknown_process"]},"allowed":false}}
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

- `event.module: agentguard`
- `event.dataset: agentguard.findings`
- `observer.vendor: AgentGuard`
- `observer.product: AgentGuard`
- `observer.type: endpoint`

## Required `agentguard.*` fields

Every finding must include:

- `agentguard.finding.id`: stable finding identifier.
- `agentguard.finding.type`: one supported finding type from this contract.
- `agentguard.finding.name`: human-readable finding name.
- `agentguard.finding.status`: `open`, `acknowledged`, or `closed`.
- `agentguard.finding.confidence`: numeric confidence from `0.0` to `1.0`.
- `agentguard.risk.level`: `low`, `medium`, `high`, or `critical`.
- `agentguard.risk.score`: integer or float from `0` to `100`.
- `agentguard.risk.reasons`: array of normalized reason labels.
- `agentguard.allowed`: boolean allowlist / approved-policy context.

Recommended lifecycle fields:

- `agentguard.finding.first_seen`
- `agentguard.finding.last_seen`
- `agentguard.allowed_by`
- `agentguard.allowed_at`

Type-specific field groups:

- AI/network: `agentguard.ai.provider`, `agentguard.ai.endpoint`, `agentguard.ai.local_service`, `agentguard.ai.model_hint`.
- MCP: `agentguard.mcp.client.name`, `agentguard.mcp.server.name`, `agentguard.mcp.server.command`, `agentguard.mcp.server.args`, `agentguard.mcp.config.path`, `agentguard.mcp.tools`, `agentguard.mcp.capabilities`.
- Browser extension: `agentguard.browser.name`, `agentguard.browser.profile`, `agentguard.extension.id`, `agentguard.extension.name`, `agentguard.extension.version`, `agentguard.extension.permissions`, `agentguard.extension.host_permissions`.
- Startup: `agentguard.startup.type`, `agentguard.startup.name`, `agentguard.startup.path`, `agentguard.startup.command`, `agentguard.startup.enabled`.
- Cyber-agent: `agentguard.cyber_agent.name`, `agentguard.cyber_agent.framework`, `agentguard.cyber_agent.provider`, `agentguard.cyber_agent.activity_type`, `agentguard.cyber_agent.capabilities`, `agentguard.cyber_agent.target_paths`, `agentguard.cyber_agent.security_tools`, `agentguard.cyber_agent.codebase_scan_volume`, `agentguard.cyber_agent.suspicious_keywords`.

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

- Producers may add new optional `agentguard.*` fields without breaking the integration when they are metadata-only and compatible with Elastic mappings.
- Producers must not change the meaning or type of existing fields.
- Risk levels must remain `low`, `medium`, `high`, or `critical`.
- `agentguard.risk.score` must remain numeric and in the `0-100` range.
- Arrays should stay arrays even when only one value is present, for example `agentguard.risk.reasons` and `agentguard.mcp.capabilities`.
- Breaking schema changes require a new contract document version.

## Privacy requirements

The producer must emit only metadata needed for detection, triage, dashboard placeholders, and draft rule references. It must not emit:

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

Use `agentguard.finding.status` consistently:

- `open`: current finding requiring attention.
- `acknowledged`: reviewed but still present.
- `closed`: no longer present or accepted through policy.

Use `agentguard.allowed: true` only when the finding matches an approved allowlist or policy exception. Include `agentguard.allowed_by` and `agentguard.allowed_at` when available.
