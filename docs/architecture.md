# AgentGuard Architecture (Current + Near-term)

## Purpose

This document describes the current module architecture in this repository and the practical direction for near-term evolution.

## Current architecture

## 1) Producer: `agentguard-sensor/`

- Runs on endpoints.
- Scans for AI-related metadata signals (for example: local services, process/context indicators, extension/config artifacts, startup/persistence-adjacent indicators).
- Emits structured findings as NDJSON (one JSON object per line).

## 2) Event shape: structured NDJSON findings

- Findings are emitted as machine-readable event records.
- ECS-compatible fields and `ai_sentinel.*` custom fields are used for downstream normalization.
- Contract and schema alignment are currently evolving and partially documented in module-specific docs.

## 3) Consumer/backend module: `elastic-integration-ai-sentinel/`

- Ingests findings via Elastic Agent collection patterns (currently file-based ingestion).
- Applies ingest pipeline parsing, ECS alignment, and privacy guardrails.
- Publishes integration assets (field definitions, pipelines, test fixtures, package metadata).

## 4) Observability workflow

1. Sensor writes NDJSON findings.
2. Elastic integration consumes records.
3. Findings become searchable/visualizable data stream events.
4. Dashboards and detection content support analyst triage and monitoring.

## Sidecar alignment direction (current)

The current repository direction includes alignment with Elastic Agent sidecar/managed deployment models:

- short term: keep file-based producer -> collector flow stable
- medium term: improve deployment and operational alignment between sensor configuration and Elastic Agent lifecycle
- longer term: evaluate tighter managed execution patterns without collapsing producer/consumer responsibilities

## Current limitations

- Backend integration is currently Elastic-first; no parity integrations are implemented here.
- Cross-session timeline correlation is limited.
- Policy/governance event layers are still draft-level.
- Browser AI visibility is early and metadata-limited.
- Shared contract versioning exists in pieces and needs stronger central references.

## Future direction

- Browser AI visibility expansion (still metadata-first).
- Timeline/session correlation across findings.
- Policy/risk event stream maturation.
- Shared producer/consumer contracts at repository root.
- Additional backend integration modules beyond Elastic.
