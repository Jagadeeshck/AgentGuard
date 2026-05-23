# AgentGuard / AI Sentinel Elastic integration

This repository hosts the `elastic-integration-ai-sentinel` project: a defensive Elastic integration package for ingesting AgentGuard / AI Sentinel NDJSON findings into Elastic Security.

> **Important:** this integration currently collects AgentGuard findings. It does not itself scan the endpoint. Active scanning is planned as the separate AgentGuard Sensor component.

## Deployment modes

This package can be deployed through:

- Fleet-managed Elastic Agent
- Standalone Elastic Agent

In both modes, Elastic Agent uses the `filestream` input to read ECS-compatible NDJSON findings written by AgentGuard / AI Sentinel and sends them to the `logs-ai_sentinel.findings-default` data stream through the package ingest pipeline.

## What this repository is

- An Elastic integration package with Fleet manifests, field definitions, ingest pipeline assets, development dashboard/rule references, sample events, and `elastic-package` tests.
- A package that reads already-produced AgentGuard / AI Sentinel finding logs from configured file paths through Elastic Agent `filestream`.
- A schema and normalization layer for ECS-compatible metadata under `ai_sentinel.*`.

## What this repository is not

- It is **not** the AgentGuard endpoint scanner.
- It does **not** add endpoint scanning code or perform host inspection.
- It does **not** add process scanning code, browser inspection code, or network capture code.
- It does **not** collect prompt content, secrets, decrypted traffic, clipboard content, or browsing history.

The AgentGuard / AI Sentinel endpoint product is a separate producer that writes the NDJSON findings consumed by this integration. See [`elastic-integration-ai-sentinel/docs/README.md`](elastic-integration-ai-sentinel/docs/README.md) for Fleet-managed and standalone installation documentation, field mapping, dashboard/rule references, privacy model, validation-pack documentation, and local `elastic-package` testing instructions. The event schema is documented in [`elastic-integration-ai-sentinel/docs/event-schema-v0.1.md`](elastic-integration-ai-sentinel/docs/event-schema-v0.1.md), and the producer contract is documented in [`elastic-integration-ai-sentinel/docs/agentguard-to-elastic-contract-v0.1.md`](elastic-integration-ai-sentinel/docs/agentguard-to-elastic-contract-v0.1.md).

## v0.3.0 validation pack

Version 0.3.0 adds validation assets, synthetic test data, a detection rule test matrix, Fleet-managed and standalone Elastic Agent deployment documentation, and the scanner-to-Elastic producer contract. This repository remains the Elastic integration only; the endpoint scanner is a separate future project. The package ingests NDJSON findings and does not include endpoint scanner, process enumeration, browser inspection, network capture, prompt collection, clipboard collection, browsing history collection, traffic decryption, or secret storage logic.
