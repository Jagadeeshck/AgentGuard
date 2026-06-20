# AgentGuard Sensor (MVP)

AgentGuard Sensor is an endpoint-side visibility scanner for AI-related activity. It emits ECS-compatible NDJSON findings using the canonical AgentGuard finding contract.

## What it does
- Scans MCP configs, browser extension manifests, local AI service ports, startup items, and AI-related processes.
- Emits one JSON event per line (`event.module=agentguard`, `event.dataset=agentguard.findings`).

## Usage
- `agentguard-sensor scan --output findings.ndjson`
- `agentguard-sensor scan --stdout`
- `agentguard-sensor scan --pretty`
- `agentguard-sensor scan --allowlist allowlist.yml`
- `agentguard-sensor version`

## Elastic integration compatibility
Output is designed for `logs-agentguard.findings-default` and compatible with the AgentGuard Elastic package field layout.

## Watch mode

```bash
agentguard-sensor watch --output findings.ndjson
agentguard-sensor watch --stdout --once
agentguard-sensor watch --interval 30s --allowlist allowlist.yml --state-file .agentguard-state.json
```

See `docs/watch-mode.md`.
