# CI Workflows

## Required workflow

The only required Elastic integration workflow is:

- **AgentGuard Elastic Integration CI** (`.github/workflows/ci.yml`)

It runs on pull requests and on pushes to `main`, and validates this repository's Elastic package with:

- `elastic-package lint`
- `elastic-package build`
- `elastic-package test pipeline`

Pipeline tests run after starting Elasticsearch only:

- `elastic-package stack up -d --services=elasticsearch`

The workflow performs guarded cleanup:

- `elastic-package stack down` only when the stack-start step succeeded and set `ELASTIC_STACK_STARTED=true`.

## Optional/manual workflows

These workflows are manual only (`workflow_dispatch`) and are not required checks.

- **AgentGuard Elastic Integration Asset Tests (Manual)** (`.github/workflows/asset-tests.yml`)
  - Runs `elastic-package test asset`
  - Starts/stops full stack for asset test execution

- **AgentGuard E2E Validation (Manual)** (`.github/workflows/e2e-agentguard.yml`)
  - Runs `scripts/e2e-validate-sensor-output.sh` only when present
  - Prints `E2E script not present; skipping.` when absent

## Sensor CI applicability

Sensor Go CI is only enabled because this repository contains Go sensor code under `agentguard-sensor/`.

- Workflow: `.github/workflows/sensor-ci.yml`
- Triggers only when files in `agentguard-sensor/**` change (plus workflow file updates)

If Go module files are not present in the repository, sensor CI should not be required.

## Local reproduction (required CI checks)

From the package directory:

```bash
cd elastic-integration-ai-sentinel
elastic-package lint
elastic-package build
elastic-package stack up -d --services=elasticsearch
elastic-package test pipeline
elastic-package stack down
```


## Package-spec hygiene

Development-only assets must live outside the package directory at `repo-root/dev-assets/elastic-integration-ai-sentinel/`.
Do not add `dev-assets/` inside `elastic-integration-ai-sentinel/`; this violates Elastic package-spec directory rules.
