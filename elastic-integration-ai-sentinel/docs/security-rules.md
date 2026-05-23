# AgentGuard Elastic Security Detection Rules

This document defines production-style Elastic Security detection rules for AgentGuard / AI Sentinel findings.

## Packaging status

As of this milestone, these rules are committed as **draft rule definitions** in `elastic-integration-ai-sentinel/dev-assets/security_rules/`.

They are not yet added under `kibana/security_rule/` because Elastic package security rule saved objects must match strict package-compatible saved object formatting. A follow-up task should convert these drafts into validated package assets once formatting is confirmed with asset validation.

## Common rule metadata

All rules share the following baseline metadata:

- `type`: `query`
- `language`: `kuery`
- `index`: `["logs-ai_sentinel.findings-*"]`
- `interval`: `5m`
- `from`: `now-10m`
- `enabled`: `false`
- `author`: `["AgentGuard"]`
- `tags`:
  - `AgentGuard`
  - `AI Sentinel`
  - `AI Activity`
  - `MCP`
  - `Elastic Security`
- `references`:
  - `https://github.com/Jagadeeshck/elastic-integration-ai-sentinel#readme`

## Rule definitions

### 1) AgentGuard Critical Finding

- **Purpose**: Detects any critical AgentGuard / AI Sentinel finding.
- **KQL**: `event.module: "ai_sentinel" and ai_sentinel.risk.level: "critical"`
- **Severity**: `critical`
- **Risk score**: `90`
- **Expected false positives**: High-risk but approved internal red-team or security engineering simulations.
- **Tuning guidance**: Mark known-approved workflows with `ai_sentinel.allowed: true` to suppress expected events.
- **Sample event type**: `ai_sentinel.risk.level: "critical"`

### 2) Untrusted MCP Server With Shell or Filesystem Access

- **Purpose**: Detects untrusted MCP servers with potentially dangerous host interaction capability.
- **KQL**: `event.module: "ai_sentinel" and ai_sentinel.finding.type: "mcp_server" and not ai_sentinel.allowed: true and ai_sentinel.mcp.capabilities: ("shell" or "filesystem")`
- **Severity**: `high`
- **Risk score**: `80`
- **Expected false positives**: Known MCP servers used by approved internal AI tooling.
- **Tuning guidance**: Allowlist known MCP config paths and set trusted MCP findings to `ai_sentinel.allowed: true`.
- **Sample event type**: `ai_sentinel.finding.type: "mcp_server"`

### 3) AI Agent With Shell Tool Use

- **Purpose**: Detects untrusted AI agent shell execution behavior.
- **KQL**: `event.module: "ai_sentinel" and ai_sentinel.finding.type: "ai_agent_shell_tool_use" and not ai_sentinel.allowed: true`
- **Severity**: `high`
- **Risk score**: `82`
- **Expected false positives**: Approved automation agents that legitimately execute shell commands.
- **Tuning guidance**: Allowlist approved process executable paths and automation identities; mark expected behavior with `ai_sentinel.allowed: true`.
- **Sample event type**: `ai_sentinel.finding.type: "ai_agent_shell_tool_use"`

### 4) Browser AI Extension With Broad Permissions

- **Purpose**: Detects AI browser extensions with broad or sensitive permissions.
- **KQL**: `event.module: "ai_sentinel" and ai_sentinel.finding.type: "browser_extension" and not ai_sentinel.allowed: true and ai_sentinel.extension.permissions: ("<all_urls>" or "scripting" or "webRequest" or "clipboardRead" or "nativeMessaging")`
- **Severity**: `medium`
- **Risk score**: `65`
- **Expected false positives**: Approved browser extensions required for enterprise workflows.
- **Tuning guidance**: Allowlist known extension IDs and set approved extensions as allowed.
- **Sample event type**: `ai_sentinel.finding.type: "browser_extension"`

### 5) Local LLM Service Exposed Beyond Loopback

- **Purpose**: Detects local AI/LLM services exposed beyond loopback interfaces.
- **KQL**: `event.module: "ai_sentinel" and ai_sentinel.finding.type: "local_llm_service" and ai_sentinel.ai.local_service: true and not destination.ip: ("127.0.0.1" or "::1" or "localhost")`
- **Severity**: `high`
- **Risk score**: `78`
- **Expected false positives**: Intentional shared local inference services in lab networks.
- **Tuning guidance**: Allowlist approved local LLM ports and known safe non-loopback exposure scenarios (for example controlled subnet service endpoints).
- **Sample event type**: `ai_sentinel.finding.type: "local_llm_service"`

### 6) AI Tool Added to Startup

- **Purpose**: Detects untrusted persistence behavior via startup registration.
- **KQL**: `event.module: "ai_sentinel" and ai_sentinel.finding.type: "startup_item" and not ai_sentinel.allowed: true`
- **Severity**: `medium`
- **Risk score**: `60`
- **Expected false positives**: Approved internal AI tooling installed for managed persistence.
- **Tuning guidance**: Allowlist approved process executable paths and startup locations; mark known-good with `ai_sentinel.allowed: true`.
- **Sample event type**: `ai_sentinel.finding.type: "startup_item"`

### 7) Possible Mythos-like AI Cyber-Agent Activity

- **Purpose**: Detects high-risk AI cyber-agent behaviors associated with vulnerability research, exploit development, or sandbox-escape style activity.
- **KQL**: `event.module: "ai_sentinel" and ai_sentinel.finding.type: ("ai_cyber_agent_activity" or "ai_vulnerability_research_agent" or "ai_exploit_development_activity" or "ai_sandbox_escape_research") and not ai_sentinel.allowed: true and ai_sentinel.risk.level: ("high" or "critical")`
- **Severity**: `critical`
- **Risk score**: `92`
- **Expected false positives**: Approved security researcher activity and sanctioned purple-team exercises.
- **Tuning guidance**: Mark sanctioned research with `ai_sentinel.allowed: true` and scope additional environment-based allowlists.
- **Sample event type**: `ai_sentinel.finding.type: "ai_cyber_agent_activity"`

### 8) AI Agent Scanning Sensitive Repositories

- **Purpose**: Detects untrusted AI agent scanning of sensitive repositories.
- **KQL**: `event.module: "ai_sentinel" and ai_sentinel.finding.type: "ai_agent_sensitive_repo_scan" and not ai_sentinel.allowed: true`
- **Severity**: `high`
- **Risk score**: `84`
- **Expected false positives**: Approved code indexing and SAST automation.
- **Tuning guidance**: Allowlist approved automation and repository scanners; mark expected scans with `ai_sentinel.allowed: true`.
- **Sample event type**: `ai_sentinel.finding.type: "ai_agent_sensitive_repo_scan"`

### 9) AI Agent Running Security Tools

- **Purpose**: Detects untrusted AI agents orchestrating security tooling.
- **KQL**: `event.module: "ai_sentinel" and ai_sentinel.cyber_agent.security_tools: * and not ai_sentinel.allowed: true`
- **Severity**: `high`
- **Risk score**: `85`
- **Expected false positives**: Approved security researcher activity and sanctioned internal assessments.
- **Tuning guidance**: Maintain allowlists for approved security toolchains and automation identities.
- **Sample event type**: `ai_sentinel.cyber_agent.security_tools: *`

### 10) AgentGuard Pipeline Failure

- **Purpose**: Detects failures in the AgentGuard ingest pipeline so SOC teams know telemetry may be degraded.
- **KQL**: `event.module: "ai_sentinel" and tags: "ai_sentinel_pipeline_failure"`
- **Severity**: `medium`
- **Risk score**: `50`
- **Expected false positives**: Short-lived deployment windows or controlled pipeline tests.
- **Tuning guidance**: Suppress during approved maintenance windows and ensure known test events are allowlisted.
- **Sample event type**: `tags: "ai_sentinel_pipeline_failure"`

## MITRE ATT&CK tactic mapping (broad)

To avoid overclaiming, mapping is limited to broad tactics only where relevant:

- **Discovery**: rules 2, 7, 8, 9
- **Execution**: rules 2, 3, 7, 9
- **Persistence**: rule 6
- **Command and Control**: rules 2, 5
- **Collection**: rules 8, 9

## Global false-positive guidance

- Approved internal AI tooling.
- Approved security researcher activity.
- Known MCP servers.
- Local-only Ollama service.
- Approved browser extensions.
- Allowlisted automation.

## Global tuning guidance

- Use `ai_sentinel.allowed: true` for known-good activity.
- Allowlist known extension IDs.
- Allowlist known MCP config paths.
- Allowlist approved process executable paths.
- Allowlist approved local LLM ports.

## TODO

Convert `dev-assets/security_rules/*.json` into package-compatible saved-object assets under `kibana/security_rule/` once asset-test-compatible formatting is validated against Elastic package security rule requirements.
