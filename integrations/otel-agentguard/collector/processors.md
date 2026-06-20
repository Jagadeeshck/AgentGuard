# AgentGuard OTEL processors

The gateway uses portable OpenTelemetry Collector processors:

- `memory_limiter` protects the gateway from unbounded queues.
- `resource/agentguard` enriches resources with deployment, organisation, and AgentGuard observer identity.
- `attributes/agentguard_defaults` normalizes logs to the AgentGuard namespace and enforces metadata-only privacy defaults.
- `batch` reduces backend load and works with Elastic, OpenSearch, Splunk, Datadog, file/NDJSON, and generic OTLP exporters.

The config intentionally avoids backend-specific processors. Optional exporters may be enabled per deployment, but the normalization model remains `agentguard.*`.
