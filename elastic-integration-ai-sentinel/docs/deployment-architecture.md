# Deployment architecture

The AgentGuard / AI Sentinel endpoint scanner is a separate product. This repository contains only the Elastic integration package that configures Elastic Agent to read scanner-produced NDJSON findings and sends them through the package ingest pipeline.

## Fleet-managed mode

```text
AgentGuard Scanner
  ↓ writes NDJSON
Elastic Agent
  ↓ filestream
Elastic ingest pipeline
  ↓
logs-ai_sentinel.findings-default
  ↓
Elastic Security
```

Fleet installs package assets, renders the `filestream` input from the integration variables, and deploys that input to enrolled Elastic Agents.

## Standalone mode

```text
AgentGuard Scanner
  ↓ writes NDJSON
Standalone Elastic Agent
  ↓ filestream input from elastic-agent.yml
Elasticsearch ingest pipeline
  ↓
logs-ai_sentinel.findings-default
```

Standalone mode uses an explicit `elastic-agent.yml`. You must install package assets and reference the ingest pipeline yourself.
