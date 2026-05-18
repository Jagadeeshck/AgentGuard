# Local Elastic Package Test Lab

This lab validates the AI Sentinel Elastic integration without requiring the endpoint scanner. It uses checked-in NDJSON fixtures, the ingest pipeline, development Kibana references, and `elastic-package` commands only.

## Scope and safety boundaries

This repository remains the Elastic integration package only. The lab does not scan endpoints, inspect local files, decrypt traffic, scrape browser history, read clipboard data, or collect prompt contents or secrets. Test data must be synthetic and limited to metadata that the package contract allows.

Allowed validation inputs:

- Synthetic NDJSON finding events in `repo-root/dev-assets/sample_events/sample_events.ndjson`.
- Pipeline fixtures in `data_stream/findings/_dev/test/pipeline`.
- Package manifests and fields inside the package, with development dashboard/rule references kept under `repo-root/dev-assets/`.

Prohibited validation inputs:

- Real prompt or completion text.
- API keys, OAuth tokens, cookies, session identifiers, or private keys.
- Clipboard content, browser history, decrypted payloads, or private file contents.
- Endpoint scanner output gathered from a real host without an explicit producer-side privacy review.

## Prerequisites

Install Docker and either install `elastic-package` locally or run commands through the containerized `ELASTIC_PACKAGE` override used by CI.

Local binary example:

```bash
cd elastic-integration-ai-sentinel
make lint
make build
make test-pipeline
```

Containerized example:

```bash
cd elastic-integration-ai-sentinel
export ELASTIC_PACKAGE='docker run --rm -v "$PWD:/package" -v /var/run/docker.sock:/var/run/docker.sock -w /package docker.elastic.co/elastic-package/elastic-package:latest'
make lint
make build
make test-pipeline
```

## Validation workflow

1. Run `make lint` to validate package structure, manifests, field definitions, and static assets.
2. Run `make build` to build the package artifact and catch packaging regressions.
3. Run `make test-pipeline` to parse synthetic NDJSON events through the ingest pipeline.
4. Run `make stack-up` before asset tests that require a local Elastic Stack.
5. Run `make test-asset` to validate installable assets against the running stack.
6. Run `make stack-down` when finished.

## Sample event replay

Use `repo-root/dev-assets/sample_events/sample_events.ndjson` as the canonical synthetic corpus for demos, manual ingest testing, and producer contract review. Each line represents a complete finding type and intentionally avoids secret values, prompt content, clipboard data, browsing history, decrypted traffic, and private file contents.

To inspect finding coverage locally:

```bash
jq -r '.ai_sentinel.finding.type' repo-root/dev-assets/sample_events/sample_events.ndjson | sort | uniq -c
```

To create a temporary pipeline fixture from a sample event, copy a single NDJSON line into `data_stream/findings/_dev/test/pipeline/test-<case>.log` and create a matching `test-<case>.log-expected.json` document following existing fixtures. Keep test values synthetic and metadata-only.

## Expected success criteria

A successful local lab run demonstrates that the integration is independently testable before an endpoint scanner exists:

- The package lints successfully.
- The package builds successfully.
- Pipeline tests pass with synthetic events.
- Asset tests pass against a local stack when Docker resources are available.
- The sample corpus covers every documented finding type.
