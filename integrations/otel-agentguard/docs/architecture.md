# Architecture

AgentGuard OTEL observes AI metadata through source adapters, emits OTEL logs/metrics/traces, normalizes into `agentguard.*`, and exports to one or more backends.

```text
AI source -> AgentGuard source adapter -> OTEL -> AgentGuard OTEL Gateway -> normalization -> backends
```

## Source adapters

Adapters convert native OTEL tools, API gateways, browser metadata, MCP/runtime events, endpoint sensor output, and SaaS audit exports into the AgentGuard semantic model.

## OTEL gateway and normalization

The gateway receives OTLP on gRPC `4317` and HTTP `4318`, applies memory limiting, resource enrichment, default privacy attributes, batching, and backend-neutral attribute normalization.

## Backend exporters

Elastic, OpenSearch, Splunk, Datadog, file/NDJSON, and generic OTLP are exporter choices. Elastic assets in this module are optional examples only.

## Endpoint sensor correlation

Existing `agentguard-sensor` discovery can be converted to OTEL events or correlated by host, user hash, process, network destination, session, and time window. The sensor path remains separate and should not be broken by this module.
