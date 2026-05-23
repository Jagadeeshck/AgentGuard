# Elastic E2E validation for AgentGuard Sensor output

Generate deterministic synthetic findings:

```bash
agentguard-sensor generate-test-findings --output findings.ndjson
```

Validate the generated NDJSON locally:

```bash
agentguard-sensor validate-output --input findings.ndjson
```

The generated file contains synthetic `ai_sentinel.*` findings for:

- `mcp_server`
- `browser_extension`
- `local_llm_service`
- `startup_item`
- `suspicious_agent_process`

Elastic Agent collects `findings.ndjson` via `filestream`, and the AI Sentinel package ingest pipeline parses, normalizes, and maps the findings into the default data stream:

- `logs-ai_sentinel.findings-default`
