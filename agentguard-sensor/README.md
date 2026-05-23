# AgentGuard Sensor (MVP)

AgentGuard Sensor is an endpoint-side defensive visibility scanner for AI-related activity. It emits ECS-compatible NDJSON findings for the Elastic `ai_sentinel` integration.

## What it does
- Scans MCP configs, browser extension manifests, local AI service ports, startup items, and suspicious AI-related processes.
- Emits one JSON event per line (`event.module=ai_sentinel`, `event.dataset=ai_sentinel.findings`).

## What it does NOT do
No prompt capture, clipboard capture, browser history collection, private document collection, traffic decryption, secret exfiltration, exploitation, persistence, or stealth behavior.

## Privacy model
Data is metadata-only and secret-like values are redacted.

## Usage
- `agentguard-sensor scan --output findings.ndjson`
- `agentguard-sensor scan --stdout`
- `agentguard-sensor scan --pretty`
- `agentguard-sensor scan --allowlist allowlist.yml`
- `agentguard-sensor version`

## Elastic integration compatibility
Output is designed for `logs-ai_sentinel.findings-default` and compatible with the `elastic-integration-ai-sentinel` package field layout.

## Example NDJSON event
```json
{"@timestamp":"2026-01-01T00:00:00Z","ecs":{"version":"8.11.0"},"event":{"module":"ai_sentinel","dataset":"ai_sentinel.findings","kind":"alert","category":["configuration"],"type":["info"],"action":"mcp_server","outcome":"success","risk_score":40,"severity":40},"observer":{"vendor":"AgentGuard","product":"AgentGuard Sensor","type":"endpoint"},"host":{"name":"host1"},"ai_sentinel":{"finding":{"id":"abc","type":"mcp_server","name":"shell server","status":"open","confidence":0.7},"risk":{"level":"medium","score":40,"reasons":["mcp shell capability"]},"allowed":false}}
```

## Uninstall/cleanup
Delete the binary and findings file path.


## Watch mode

```bash
agentguard-sensor watch --output findings.ndjson
agentguard-sensor watch --stdout --once
agentguard-sensor watch --interval 30s --allowlist allowlist.yml --state-file .agentguard-state.json
```

See `docs/watch-mode.md`.
