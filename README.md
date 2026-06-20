# AgentGuard

AgentGuard is a full **AI activity observability platform**. It provides endpoint-side visibility into AI agent activity, structured findings output, and a pluggable backend integration layer. Elastic is the first supported backend, not the product itself.

This repository contains:

- `agentguard-sensor/` — endpoint-side producer that discovers AI-related runtime signals
- Elastic integration assets — first backend integration for Elastic Agent and Elastic Security ingestion
- `dev-assets/` — development assets, examples, and test data

## Overview

AgentGuard currently implements an observability-first MVP:

1. The sensor module discovers AI-related runtime and configuration signals on endpoints.
2. Findings are emitted as structured AgentGuard NDJSON events.
3. Backend integration modules ingest, normalize, and operationalize those findings.
4. Security and operations teams investigate in their SIEM/observability tooling.

The current backend path is Elastic.

## Current status

- Endpoint producer logic is under active development in `agentguard-sensor/`.
- Elastic integration package assets are maintained for AgentGuard findings.
- CI currently validates both sensor (Go tests/build) and Elastic package workflows.
- Sidecar-aligned deployment direction is documented and in progress for Elastic Agent integration patterns.

## Repository modules

### `agentguard-sensor/`
Endpoint-side scanner/sensor component that emits AgentGuard findings as NDJSON events designed for downstream ingestion.

### Elastic integration assets
Elastic package module that consumes AgentGuard sensor output using Elastic Agent collection patterns and maps findings into ECS-aligned fields, ingest pipelines, and package assets.

### `dev-assets/`
Supporting assets such as placeholder dashboards, detection-rule drafts, examples, synthetic test data, and validation helpers.

## Architecture summary

- **Producer:** `agentguard-sensor/`
- **Transport/data shape:** file-based AgentGuard NDJSON structured findings
- **Current consumer/backend module:** Elastic integration assets
- **Analyst workflow:** ingest -> normalize -> search/dashboard/detection

This separation keeps producer behavior independent from backend-specific integration logic.

## Privacy model

AgentGuard follows a metadata-first privacy posture in the current implementation.

See `docs/privacy-model.md` for the current model and guardrails.

## Current MVP scope

In-scope today:

- endpoint metadata collection for AI-related signals
- structured findings output
- Elastic ingestion and normalization path
- draft detection/dashboard assets for iterative validation

Out-of-scope today:

- full browser activity instrumentation beyond currently implemented metadata sources
- deep timeline/session correlation across all event classes
- broad multi-backend parity

## Roadmap themes

- stabilize producer/consumer contracts
- harden sensor coverage and event quality
- mature Elastic package dashboards and detection workflows
- add browser AI visibility incrementally
- add policy/risk and correlation layers
- add additional backend/export integrations

See `docs/roadmap.md` and `docs/mvp-backlog.md`.

## Current structure

```text
.
├── agentguard-sensor/
├── Elastic integration assets/
├── dev-assets/
├── .github/workflows/
├── contracts/
└── docs/
```

## Recommended future structure

A future evolution can move backend-specific modules under a neutral `integrations/` directory:

```text
.
├── agentguard-sensor/
├── integrations/
│   └── elastic/
├── contracts/
├── docs/
├── dev-assets/
└── .github/workflows/
```

The Elastic backend remains the first supported integration module for AgentGuard findings.
