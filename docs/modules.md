# AgentGuard Repository Modules

## Current modules

## `agentguard-sensor/`

Endpoint-side producer module.

- Role: discovery/scanning and structured finding emission.
- Output: NDJSON findings intended for backend ingestion.
- Ownership boundary: endpoint collection behavior and producer-side event quality.

## `integrations/elastic-agentguard/`

Current backend integration module (Elastic-first).

- Role: ingest and normalize findings into Elastic via package assets.
- Includes: ingest pipelines, fields, package manifests, tests, and module docs.
- Ownership boundary: Elastic-specific mapping, packaging, and operational guidance.

## `dev-assets/`

Development-only supporting assets.

- Includes sample events, placeholders, examples, and rule/dash drafts.
- Not a production runtime module by itself.

## `.github/workflows/`

Automation boundary for CI/validation.

- Includes sensor CI and Elastic package validation/test workflows.
- Enforces non-regression for module-specific pipelines.

## Current module boundaries

- Sensor and backend integration are separate modules with explicit producer/consumer responsibilities.
- Elastic integration must not be treated as endpoint scanning logic.
- Shared documentation/contracts should live at root and complement module docs.

## Future planned modules

- Additional backend integrations/exports.
- Shared validation utilities for producer-consumer contract checks.
- Optional correlation/policy processing components as separate modules if needed.

## Recommended future repo layout

```text
.
├── agentguard-sensor/
├── integrations/elastic-agentguard/
├── integrations/                 # future backend modules (recommended)
├── contracts/                    # shared contracts and schemas
├── docs/                         # product-level architecture and roadmap docs
├── dev-assets/
└── .github/workflows/
```

The current `integrations/elastic-agentguard/` folder name remains unchanged in this repository state.
