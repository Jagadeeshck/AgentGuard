# AgentGuard

AgentGuard is a full **AI activity observability platform** — not just an Elastic integration. It provides endpoint-side visibility into AI agent activity, structured findings output, and a pluggable backend integration layer. Elastic is the first supported backend, not the product itself.

This repository contains:

- `agentguard-sensor/` — endpoint-side producer that discovers AI-related runtime signals
- `elastic-integration-ai-sentinel/` — first backend integration (Elastic SIEM)
- `dev-assets/` — development assets, examples, and test data

## Overview

AgentGuard currently implements an observability-first MVP:

1. The sensor module discovers AI-related runtime and configuration signals on endpoints.
2. Findings are emitted as structured NDJSON events.
3. Backend integration modules ingest, normalize, and operationalize those findings.
4. Security and operations teams investigate in their SIEM/observability tooling.

The current backend path is Elastic via the `elastic-integration-ai-sentinel` module.

## Current status

- Endpoint producer logic is under active development in `agentguard-sensor/`.
- Elastic integration package assets are maintained in `elastic-integration-ai-sentinel/`.
- CI currently validates both sensor (Go tests/build) and Elastic package workflows.
- Sidecar-aligned deployment direction is documented and in progress for Elastic Agent integration patterns.

## Repository modules

### `agentguard-sensor/`
Endpoint-side scanner/sensor component that emits findings as NDJSON events designed for downstream ingestion.

### `elastic-integration-ai-sentinel/`
Elastic package module that consumes sensor output using Elastic Agent collection patterns (currently `filestream`-based) and maps findings into ECS-aligned fields, ingest pipelines, and package assets.

### `dev-assets/`
Supporting assets such as placeholder dashboards, detection-rule drafts, examples, synthetic test data, and validation helpers.

## Architecture summary

- **Producer:** `agentguard-sensor/`
- **Transport/data shape:** file-based NDJSON structured findings
- **Current consumer/backend module:** `elastic-integration-ai-sentinel/`
- **Analyst workflow:** ingest -> normalize -> search/dashboard/detection

This separation keeps producer behavior independent from backend-specific integration logic.

## Privacy model

AgentGuard follows a metadata-first privacy posture in the current implementation:

- no prompt capture by default
- no clipboard capture by default
- no browsing history collection by default
- no decrypted traffic inspection
- no secret/credential capture intent

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
├── elastic-integration-ai-sentinel/
├── dev-assets/
├── .github/workflows/
├── contracts/
└── docs/
```

## Recommended future structure

No top-level directories are renamed in this change. A future evolution can add module growth without breaking existing paths:

```text
.
├── agentguard-sensor/
├── elastic-integration-ai-sentinel/
├── integrations/
│   └── <future-backend-modules>
├── contracts/
├── docs/
├── dev-assets/
└── .github/workflows/
```

The `elastic-integration-ai-sentinel` name remains unchanged and is treated as the first backend integration module.
