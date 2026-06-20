# Backend exporters

AgentGuard OTEL is backend-neutral.

- Elastic: optional index template, component template, and ingest pipeline examples for `logs-agentguard.ai_activity-*`.
- OpenSearch: send OTLP through a collector exporter or write NDJSON with preserved `agentguard.*` attributes.
- Splunk: export through OTLP HTTP, HEC, or collector-managed routing.
- Datadog: export OTLP logs/metrics/traces and map AgentGuard attributes as tags/facets.
- file/NDJSON: durable local archive or replay format for testing and air-gapped workflows.
- generic OTLP: forward to any compliant backend.
