# AgentGuard

This repository hosts the `elastic-integration-ai-sentinel` project: a defensive Elastic Fleet integration package for ingesting AI Sentinel NDJSON findings into Elastic Security.

## What this repository is

- An Elastic integration package with Fleet manifests, field definitions, ingest pipeline assets, Kibana dashboards, detection rules, sample events, and `elastic-package` tests.
- A package that reads already-produced AI Sentinel finding logs from configured file paths through Elastic Agent `filestream`.
- A schema and normalization layer for ECS-compatible metadata under `ai_sentinel.*`.

## What this repository is not

- It is **not** the AI Sentinel endpoint scanner.
- It does **not** add endpoint scanning code or perform host inspection.
- It does **not** collect prompt content, secrets, decrypted traffic, clipboard content, or browsing history.

The AI Sentinel endpoint product is a separate producer that writes the NDJSON findings consumed by this integration. See [`elastic-integration-ai-sentinel/docs/README.md`](elastic-integration-ai-sentinel/docs/README.md) for package installation, field mapping, dashboards, rules, privacy model, and local `elastic-package` testing instructions. The event schema is documented in [`elastic-integration-ai-sentinel/docs/event-schema-v0.1.md`](elastic-integration-ai-sentinel/docs/event-schema-v0.1.md).
