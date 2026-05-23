# Elastic Agent-managed AgentGuard Sensor feasibility (Phase 1)

## Summary

For the current package/repo structure, **MVP support is sidecar mode**: AgentGuard Sensor runs as an OS service, and Elastic Agent collects `findings.ndjson` via `filestream`.

## Options reviewed

### Option A: Integration package directly runs a third-party binary

Not supported for this package type. An Elastic integration package by itself provides package assets (manifests, ingest pipelines, field mappings, and input templates) and cannot safely bundle arbitrary executables and process lifecycle control for Fleet.

### Option B: Integration package + Fleet-managed config + external sensor service

Supported and recommended for MVP. The integration can expose Fleet variables that align with AgentGuard Sensor configuration, while Elastic Agent continues to read NDJSON output paths.

### Option C: External OS package + Fleet-managed collection

Also supported; effectively the same deployment model as Option B with stronger separation between sensor lifecycle and log collection.

## Required metadata/files for supported MVP

- `data_stream/findings/manifest.yml`: Fleet variables and user-facing descriptions.
- `data_stream/findings/agent/stream/filestream.yml.hbs`: filestream template consuming configured paths.
- Package docs explaining sidecar deployment and verification.
- Sensor CLI/docs to generate and verify Elastic-aligned paths.

## What elastic-package validates

`elastic-package` validates package structure, manifests, data stream assets, ingest pipelines, tests, and built package artifacts. It does **not** validate that Elastic Agent can run an external third-party scanner process from this integration package.

## Safe MVP

Safe MVP is:

1. AgentGuard Sensor installed and managed as OS service.
2. Sensor writes ECS-compatible NDJSON to a configured path.
3. Fleet policy configures matching `paths` for filestream collection.
4. Elastic Agent forwards events to `logs-ai_sentinel.findings-default`.

## Recommended architecture

Current milestone recommendation:

- **Now (supported):** Fleet-managed sidecar mode.
- **Future (investigation):** native Elastic Agent custom input/component only if implemented in a dedicated Elastic Agent component project and proven runnable/maintainable.

No claim should state Elastic Agent runs the scanner unless a custom component exists and is validated.
