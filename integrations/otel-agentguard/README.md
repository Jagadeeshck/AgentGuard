# AgentGuard OpenTelemetry Integration

AgentGuard OTEL is a vendor-neutral module for observing AI activity across web AI tools, APIs, coding assistants, local LLMs, MCP servers, frameworks, SaaS audit exports, and existing AgentGuard endpoint discovery. It is **not** Claude Code monitoring and does not depend on any single backend.

Canonical values:

- Product: `AgentGuard`
- Event module: `agentguard`
- Default findings dataset: `agentguard.findings`
- Generic OTEL dataset: `agentguard.otel`
- AI activity dataset: `agentguard.ai_activity`
- Namespace: `agentguard.*`
- Privacy posture: metadata-first; prompt and content capture are disabled by default.

## Pipeline

```text
AI source
  -> AgentGuard source adapter
  -> OTEL logs/metrics/traces
  -> AgentGuard OTEL Gateway
  -> AgentGuard normalization processor
  -> one or more backends
       - Elastic
       - OpenSearch
       - Splunk
       - Datadog
       - file/NDJSON
       - generic OTLP backend
```

## Safe OpenAI and ChatGPT adapter patterns

AgentGuard does not assume direct internal ChatGPT telemetry and does not implement undocumented scraping.

- `openai_api_gateway`: metadata from internal applications, approved gateways, or API logs for OpenAI API usage.
- `chatgpt_browser_metadata`: metadata-only ChatGPT web usage through a future enterprise browser extension or managed browser logs.
- `chatgpt_enterprise_audit_placeholder`: placeholder for future approved enterprise audit export integrations.

## Privacy defaults

AgentGuard is not spyware. Deploy only on owned/managed systems or with explicit user and enterprise consent. Prompt text, chat content, clipboard content, browser history, decrypted traffic, and raw request/response bodies must not be captured by default.

Default attributes:

```text
agentguard.privacy.prompt_capture_enabled=false
agentguard.privacy.content_capture_enabled=false
agentguard.privacy.redaction_status=metadata_only
```

## Contents

- `collector/`: portable OpenTelemetry Collector gateway example.
- `schema/`: AgentGuard neutral semantic conventions and JSON Schema.
- `examples/`: synthetic metadata-only events and environment examples.
- `elastic/`: optional Elastic backend assets only.
- `detections/`: backend-neutral detection drafts represented as JSON.
- `docs/`: architecture, privacy, adapters, exporters, and roadmap.

## Collector exporters

The sample collector writes logs to both `debug` and `file/ndjson` by default so it is useful without an external backend. To send logs to another backend, add `otlphttp/generic` or `elasticsearch/optional` to `service.pipelines.logs.exporters` in `collector/otel-collector-config.yaml` and provide the matching environment variables.
