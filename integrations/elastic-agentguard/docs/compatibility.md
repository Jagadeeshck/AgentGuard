# Compatibility

## Supported Elastic Stack versions

The package manifest declares Kibana compatibility with `^8.13.0 || ^9.0.0`.

## Supported Elastic Agent versions

Use an Elastic Agent version compatible with the target Elastic Stack and Fleet version. The package uses the Elastic Agent `filestream` input and has been structured for `elastic-package` lint, build, and pipeline tests.

## Supported OS paths

Default findings paths are:

- Linux: `/var/log/agentguard/findings.ndjson`, `/var/log/agentguard/findings.ndjson`
- macOS: `/Library/Logs/AgentGuard/findings.ndjson`, `/Library/Logs/AI-Sentinel/findings.ndjson`
- Windows: `C:\ProgramData\AgentGuard\logs\findings.ndjson`, `C:\ProgramData\AI-Sentinel\logs\findings.ndjson`

## Deployment modes

- Fleet-managed Elastic Agent: supported through the package policy template and `filestream` stream template.
- Standalone Elastic Agent: supported through manual `elastic-agent.yml` configuration that uses the same dataset and ingest pipeline.

## Expected data stream

For the default namespace, events land in:

```text
logs-agentguard.findings-default
```

## Expected ingest pipeline behavior

The package ingest pipeline parses NDJSON, sets `event.module: agentguard` and `event.dataset: agentguard.findings`, adds AgentGuard tags, maps risk metadata, removes prohibited raw content fields, redacts common secret patterns, and drops `event.original` unless preservation is explicitly enabled.

## Known limitations

- This package does not include the AgentGuard endpoint scanner.
- This package does not create additional datasets yet.
- Standalone mode requires package assets and ingest pipelines to be installed separately.
- Detection quality depends on the scanner writing valid ECS-compatible NDJSON findings.
