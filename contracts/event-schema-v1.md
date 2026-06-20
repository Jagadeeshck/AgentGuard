# AgentGuard Event Schema v1 (Draft)

## Purpose

Define a practical canonical event model for AgentGuard MVP findings emitted by producers and consumed by backend integrations.

## Scope

- Applies to structured NDJSON finding events.
- Covers common required fields and category/type conventions.
- Complements backend-specific mappings, for example Elastic package details.

## Versioning

- Schema version field: `agentguard.schema.version`.
- Current draft: `1.0.0-draft`.
- Compatibility rule: additive fields are backward compatible; required-field removals are breaking.

## Required common fields

Each event must include:

- `@timestamp`
- `event.module`
- `event.dataset`
- `event.kind`
- `event.category` (array)
- `event.type` (array)
- `event.action`
- `event.outcome`
- `agentguard.schema.version`
- `agentguard.finding.id`
- `agentguard.finding.type`
- `agentguard.risk.level`
- `agentguard.risk.score`
- `observer.vendor`
- `observer.product`
- `host.name` (where available)

## Recommended ECS alignment

- Keep ECS-standard fields at ECS paths (`host.*`, `process.*`, `user.*`, `file.*`, `network.*`, `source.*`, `destination.*`).
- Keep AgentGuard custom fields under `agentguard.*`.
- Backend integrations may transform transport-specific fields while preserving AgentGuard semantics.

## Event categories and example event types

Canonical event types (initial set):

- `agentguard.finding.detected`
- `agentguard.service.detected`
- `agentguard.browser_extension.detected`
- `agentguard.process.detected`
- `agentguard.persistence.detected`
- `agentguard.policy.match`

## Sample events

### 1) Generic finding

```json
{"@timestamp":"2026-01-01T00:00:00Z","event":{"module":"agentguard","dataset":"agentguard.findings","kind":"alert","category":["configuration"],"type":["info"],"action":"finding_detected","outcome":"success"},"observer":{"vendor":"AgentGuard","product":"AgentGuard Sensor"},"host":{"name":"host-a"},"agentguard":{"schema":{"version":"1.0.0-draft"},"finding":{"id":"f-001","type":"agentguard.finding.detected","name":"mcp_server_detected"},"risk":{"level":"medium","score":45,"reasons":["mcp_shell_capability"]}}}
```

### 2) Browser extension finding

```json
{"@timestamp":"2026-01-01T00:01:00Z","event":{"module":"agentguard","dataset":"agentguard.findings","kind":"alert","category":["configuration"],"type":["change"],"action":"browser_extension_detected","outcome":"success"},"observer":{"vendor":"AgentGuard","product":"AgentGuard Sensor"},"host":{"name":"host-b"},"agentguard":{"schema":{"version":"1.0.0-draft"},"finding":{"id":"f-002","type":"agentguard.browser_extension.detected","name":"extension_permission_profile"},"risk":{"level":"high","score":75,"reasons":["broad_permissions"]},"browser_extension":{"id":"abcd","name":"Example AI Helper","permissions":["<all_urls>","nativeMessaging"]}}}
```

### 3) Policy match event

```json
{"@timestamp":"2026-01-01T00:02:00Z","event":{"module":"agentguard","dataset":"agentguard.findings","kind":"signal","category":["policy"],"type":["indicator"],"action":"policy_match","outcome":"success"},"observer":{"vendor":"AgentGuard","product":"AgentGuard Sensor"},"host":{"name":"host-c"},"agentguard":{"schema":{"version":"1.0.0-draft"},"finding":{"id":"f-003","type":"agentguard.policy.match","name":"unapproved_mcp_capability"},"risk":{"level":"high","score":82,"reasons":["policy_violation"]},"policy":{"id":"pol-001","name":"MCP capability baseline","result":"match"}}}
```

## Elastic package alignment

- Package name: `agentguard`
- Dataset: `agentguard.findings`
- Data stream: `logs-agentguard.findings-default`
- Canonical custom namespace: `agentguard.*`
