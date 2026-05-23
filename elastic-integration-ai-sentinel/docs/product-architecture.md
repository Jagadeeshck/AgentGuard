# Product architecture

This document clarifies the current and target architecture for AgentGuard / AI Sentinel.

## 1) Current package role

This repository is currently the **AgentGuard Elastic Integration** package (Elastic Agent/Fleet integration layer), not the active endpoint scanner.

It currently does the following:

- Collects `findings.ndjson` from configured paths via Elastic Agent `filestream`
- Parses JSON events
- Normalizes events to ECS-compatible fields
- Stores custom fields under `ai_sentinel.*`
- Provides a base for dashboards and detection rules
- Supports both Fleet-managed and standalone Elastic Agent collection

It currently does **not** do the following:

- Scan processes
- Inspect browser extensions
- Inspect MCP configs
- Inspect cloud accounts
- Inspect network sockets
- Monitor local LLM services
- Install an endpoint sensor
- Replace Elastic Defend

## 2) Full target architecture

The full AgentGuard architecture is intended to include three components.

### A) AgentGuard Sensor

- Endpoint/cloud scanner and monitor
- Discovers AI tools, MCP servers, browser extensions, local LLMs, startup items, AI API connections, and AI cyber-agent behavior
- Writes ECS-compatible NDJSON findings
- Future: may send directly to Elasticsearch
- Future: may be managed by Elastic Agent/Fleet if custom input packaging is implemented

### B) AgentGuard Elastic Integration

- **This repository**
- Collects sensor output
- Parses, maps, enriches, and redacts findings
- Provides dashboards and detection rules

### C) Optional AgentGuard Detection Pack

- Future package/content that uses existing Elastic Defend/System/Auditd/Sysmon/cloud telemetry
- Detects AI activity without requiring the AgentGuard Sensor where possible

## 3) Deployment modes

### Mode 1: MVP / current

- AgentGuard Sensor is installed separately
- Sensor writes `findings.ndjson`
- Elastic Agent `filestream` collects `findings.ndjson`
- Elastic Security stores data and runs alerts

### Mode 2: Future Elastic Agent-managed sensor

- Fleet installs AgentGuard integration
- Elastic Agent launches/manages AgentGuard Sensor
- Sensor emits events
- Elastic Agent ships events

### Mode 3: Elastic Defend augmentation

- Elastic Defend collects endpoint telemetry
- AgentGuard detection rules identify AI-agent/MCP/local-LLM behavior from existing endpoint events

## 4) Why this separation matters

- Keeps the Elastic package clean and package-spec compliant
- Keeps scanner development independent
- Avoids pretending `filestream` scans the host
- Enables enterprise deployment through Fleet
- Enables future integration with Elastic Defend telemetry

## 5) Roadmap

- **v0.3.x**
  - Stable ingestion package
  - Pipeline tests passing
  - NDJSON producer contract
- **v0.4**
  - Dashboards and detection rules
- **v0.5**
  - AgentGuard Sensor MVP in a separate repository
- **v0.6**
  - Elastic Defend-based detection pack
- **v1.0**
  - Optional Elastic Agent-managed sensor deployment
