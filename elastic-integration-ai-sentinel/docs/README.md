# AgentGuard / AI Sentinel Elastic Integration

AgentGuard / AI Sentinel is a defensive endpoint visibility product that emits ECS-compatible NDJSON findings about AI-related activity such as AI API connections, MCP server configuration, local LLM services, browser extensions, and startup items.

This repository contains **only** the Elastic integration package. It is not the AgentGuard endpoint scanner and it does not include scanner collection logic. The package can be deployed through Fleet-managed Elastic Agent or standalone Elastic Agent. It reads already-produced AgentGuard / AI Sentinel NDJSON findings from disk, parses them into ECS-compatible fields, and ships Elastic package assets such as fields and ingest pipelines, while keeping dashboard placeholders and draft rule references under `_dev/` until they are converted to supported package asset formats. It does not perform endpoint scanning, decrypt traffic, collect private prompt content, collect clipboard content, collect browsing history, or store secrets.

## Data streams

The first production data stream is:

- `logs-ai_sentinel.findings-default`

Future data stream names are reserved for more specialised telemetry:

- `logs-ai_sentinel.mcp-default`
- `logs-ai_sentinel.network-default`
- `logs-ai_sentinel.process-default`
- `logs-ai_sentinel.browser_extension-default`
- `logs-ai_sentinel.startup-default`

## Deployment modes

This package can be deployed through:

- Fleet-managed Elastic Agent
- Standalone Elastic Agent

It collects NDJSON findings written by AgentGuard / AI Sentinel. It does not contain the AgentGuard endpoint scanner.

## Fleet installation

See [fleet-installation.md](fleet-installation.md) for the complete Fleet-managed Elastic Agent workflow.

1. Build the package with `elastic-package build`.
2. Add the package zip from `build/packages/` to your internal Elastic Package Registry or install it in a development stack with `elastic-package stack up` and `elastic-package install`.
3. In Kibana Fleet, add the **AgentGuard / AI Sentinel** integration to an Elastic Agent policy.
4. Confirm that the AgentGuard / AI Sentinel endpoint product writes NDJSON findings to one or more configured paths.

## Log collection

The integration uses Elastic Agent `filestream` to read JSON lines from disk. Default paths are:

- Linux: `/var/log/agentguard/findings.ndjson`, `/var/log/ai-sentinel/findings.ndjson`
- macOS: `/Library/Logs/AgentGuard/findings.ndjson`, `/Library/Logs/AI-Sentinel/findings.ndjson`
- Windows: `C:\ProgramData\AgentGuard\logs\findings.ndjson`, `C:\ProgramData\AI-Sentinel\logs\findings.ndjson`

Configurable variables:

- `paths`: one or more NDJSON paths.
- `tags`: tags added by Elastic Agent.
- `preserve_original_event`: when enabled, copies the raw log line to `event.original`.
- `processors`: optional Elastic Agent processors.
- `timezone`: local timezone metadata for collectors that need it.
- `data_stream.dataset`: defaults to `ai_sentinel.findings`.
- `pipeline`: defaults to `logs-ai_sentinel.findings-default`.
- `ignore_older`, `close_inactive`, `exclude_files`, `include_files`, and `parsers`: advanced filestream settings.

## Standalone Elastic Agent

See [standalone-elastic-agent.md](standalone-elastic-agent.md) and the documentation-only examples under `repo-root/dev-assets/examples/` for standalone `elastic-agent.yml` configuration.

## Event schema

The MVP event contract is documented in [event-schema-v0.1.md](event-schema-v0.1.md). Producers should emit one JSON object per line and keep all non-ECS custom fields under `ai_sentinel.*`. Supported finding types are documented in [event-taxonomy.md](event-taxonomy.md), including base AI visibility findings and AI cyber-agent detection pack findings.

## ECS mapping model

AI Sentinel events should already be ECS-compatible. The ingest pipeline safely parses JSON, sets `event.module: ai_sentinel`, sets `event.dataset: ai_sentinel.findings`, populates observer metadata, adds AI Sentinel tags, maps `ai_sentinel.risk.score` to `event.risk_score`, derives `event.severity` from `ai_sentinel.risk.level`, and preserves ECS fields for `host`, `process`, `user`, `file`, `source`, `destination`, `network`, `dns`, and `related`.

Custom fields live under `ai_sentinel.*`. Required groups include finding identity, risk details, AI provider/endpoint metadata, MCP client and server metadata, browser extension metadata, startup item metadata, and allowlist state.

Risk levels are `low`, `medium`, `high`, and `critical`.

## Redaction and privacy

The ingest pipeline redacts common secret patterns in command lines, MCP server args, URL query strings, authorization headers, startup commands, and configuration-like paths where possible. AI Sentinel should avoid emitting private content, raw prompts, decrypted traffic, or credentials. Treat any optional `event.original` retention as sensitive and enable it only for controlled debugging.

## Dashboards

Development placeholder dashboard drafts were moved to `repo-root/dev-assets/kibana_placeholders/` for these dashboard entry points:

1. AI Sentinel Overview: findings by risk, host, time, providers, risky processes, and critical/high tables.
2. MCP Security Dashboard: MCP servers, clients, capabilities, privileged access, and changed configs.
3. AI Network Activity Dashboard: AI providers, destination domains, connecting processes, and local services.
4. Browser Extension Risk Dashboard: AI-related extensions, broad permissions, native messaging, and all-sites access.
5. Endpoint AI Inventory: AI tools by host, local LLM services, MCP-enabled clients, startup items, and allowlisted vs untrusted findings.

The placeholders define stable IDs so maintainers can replace them with production Lens panels without changing downstream references.

## Detection rules

Draft Security detection rules are listed in [security-rules.md](security-rules.md), with TOML drafts kept under `repo-root/dev-assets/security_rules_toml/` until they are converted to package-supported saved-object JSON.

Rules target `logs-ai_sentinel.findings-*` and use KQL against ECS and `ai_sentinel.*` fields. The rule coverage plan is documented in [detection-rule-test-matrix.md](detection-rule-test-matrix.md).

## Local development and testing

From the package directory:

```bash
make lint
make format
make build
make test-pipeline
make test-asset
make stack-up
make stack-down
```

Equivalent raw commands:

```bash
elastic-package lint
elastic-package format
elastic-package build
elastic-package test pipeline
elastic-package test asset
elastic-package stack up
elastic-package stack down
```

Pipeline test fixtures live in `data_stream/findings/test/pipeline/` and cover `ai_api_connection`, `mcp_server`, `browser_extension`, `startup_item`, `local_llm_service`, cyber-agent pack examples, malformed JSON, redaction, missing optional fields, risk score mapping, and event categorisation. Each `.log` fixture has a matching expected `.json` output file for `elastic-package test pipeline`. The broader synthetic validation corpus lives in `repo-root/dev-assets/sample_events/sample_events.ndjson`.

## Validation pack documentation

Version 0.3.0 adds validation, synthetic test data under `repo-root/dev-assets/test_data/`, a detection rule test matrix, and the scanner-to-Elastic contract. The endpoint scanner remains a separate future project; this package only ingests NDJSON findings.

Version 0.3.0 adds an end-to-end validation pack so this Elastic integration can be tested independently before an endpoint scanner exists:

- [Fleet Installation](fleet-installation.md)
- [Standalone Elastic Agent](standalone-elastic-agent.md)
- [Deployment Architecture](deployment-architecture.md)
- [Security and Privacy](security-and-privacy.md)
- [Compatibility](compatibility.md)
- [Local Elastic Package Lab](local-lab.md)
- [AgentGuard to Elastic Contract v0.1](agentguard-to-elastic-contract-v0.1.md)
- [Event Taxonomy](event-taxonomy.md)
- [Risk Scoring Model](risk-scoring-model.md)
- [False Positive Guidance](false-positive-guidance.md)
- [Safe-vs-Dangerous Scenarios](safe-vs-dangerous-scenarios.md)
- [Detection Rule Test Matrix](detection-rule-test-matrix.md)
- Synthetic Test Data: `repo-root/dev-assets/test_data/`

## Troubleshooting

- No events: verify the Elastic Agent policy path matches the AI Sentinel NDJSON path and that the Agent user can read the file.
- Invalid JSON tags: inspect events tagged `ai_sentinel_invalid_json`; each line must be a single valid JSON object.
- Missing dashboards or rules: placeholder dashboards and TOML rule drafts are development references only until converted to package-supported saved-object JSON assets.
- Unexpected secrets: disable `preserve_original_event`, verify AI Sentinel scanner-side redaction, and review fields that contain commands, URLs, headers, or config paths.
