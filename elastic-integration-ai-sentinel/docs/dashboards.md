# AgentGuard Dashboards

Data view for all dashboards: `logs-ai_sentinel.findings-*`.

## 1) AgentGuard AI Activity Overview

**Purpose**
- SOC-level overview of AI-related activity across endpoints.

**Panels**
- Findings over time
- Findings by risk level
- Findings by host
- Findings by type
- Top AI providers
- Top risky processes
- Critical/high findings table
- Allowlisted vs untrusted findings

**Fields used**
- `@timestamp`, `host.name`, `ai_sentinel.finding.type`, `ai_sentinel.finding.name`
- `ai_sentinel.risk.level`, `ai_sentinel.risk.score`, `ai_sentinel.allowed`
- `process.name`, `process.executable`, `ai_sentinel.ai.provider`, `event.action`

**Example questions**
- Which endpoints have untrusted AI agents?
- Are high-risk findings increasing over time?
- Which providers and processes are most represented in risky activity?

**SOC workflow**
- Start with trend and risk distributions.
- Pivot to hosts and finding types.
- Open critical/high table for triage and escalation.

**Tuning suggestions**
- Filter allowlisted findings by default for incident triage views.
- Add environment/host group filters (prod/dev) for noise reduction.

## 2) AgentGuard MCP Security Dashboard

**Purpose**
- Monitor MCP servers, exposed tools, and dangerous capabilities.

**Panels**
- MCP servers by host
- MCP clients discovered
- MCP capabilities distribution
- MCP servers with shell capability
- MCP servers with filesystem capability
- MCP servers with browser/database/code repository capability
- Recently changed MCP configs
- Untrusted high-risk MCP servers table

**Fields used**
- `ai_sentinel.mcp.client.name`, `ai_sentinel.mcp.server.name`, `ai_sentinel.mcp.server.command`
- `ai_sentinel.mcp.config.path`, `ai_sentinel.mcp.tools`, `ai_sentinel.mcp.capabilities`
- `ai_sentinel.mcp.risk_reasons`, `host.name`, `ai_sentinel.allowed`, `ai_sentinel.risk.level`

**Example questions**
- Which MCP servers expose shell or filesystem tools?
- Which hosts recently changed MCP configurations?
- Which untrusted MCP servers are high risk?

**SOC workflow**
- Review capability and tool exposure distributions.
- Focus on untrusted + high-risk table rows.
- Correlate to host and process context for response actions.

**Tuning suggestions**
- Maintain explicit allowlist policy for known-safe MCP servers.
- Alert on first-seen shell/filesystem capabilities per host.

## 3) AgentGuard Local AI / LLM Services Dashboard

**Purpose**
- Show local AI services, exposed ports, and local model tooling.

**Panels**
- Local LLM services by host
- Exposed non-loopback services
- Ports used by AI services
- Processes backing local AI services
- Local AI service risk table

**Fields used**
- `ai_sentinel.finding.type`, `ai_sentinel.ai.local_service`, `ai_sentinel.ai.provider`
- `ai_sentinel.ai.endpoint`, `destination.ip`, `destination.port`
- `process.name`, `process.executable`, `host.name`, `ai_sentinel.risk.level`

**Example questions**
- Are local LLM services exposed beyond localhost?
- Which ports are most commonly used by local AI services?
- Which processes are backing exposed local AI endpoints?

**SOC workflow**
- Identify non-loopback endpoint exposure first.
- Drill into host/process combinations and risk levels.
- Coordinate with endpoint owners for service hardening.

**Tuning suggestions**
- Suppress known approved localhost-only services.
- Flag unusual port/provider combinations for review.

## 4) AgentGuard Browser AI Extension Risk Dashboard

**Purpose**
- Find risky AI browser extensions and excessive permissions.

**Panels**
- Extensions by browser
- Extensions with all_urls access
- Extensions with nativeMessaging
- Extensions with clipboardRead
- Extensions with scripting/webRequest
- High-risk extension table

**Fields used**
- `ai_sentinel.browser.name`, `ai_sentinel.browser.profile`
- `ai_sentinel.extension.id`, `ai_sentinel.extension.name`, `ai_sentinel.extension.version`
- `ai_sentinel.extension.permissions`, `ai_sentinel.extension.host_permissions`
- `ai_sentinel.extension.risk_reasons`, `host.name`, `ai_sentinel.risk.level`

**Example questions**
- Which browser extensions have broad permissions?
- Which hosts have extensions using native messaging?
- Which high-risk extensions should be removed first?

**SOC workflow**
- Triage permissions-oriented panels first.
- Pivot into high-risk extension table and affected host set.
- Initiate extension block/remove workflow with endpoint teams.

**Tuning suggestions**
- Keep approved extension inventory current.
- Review host permissions for wildcard patterns regularly.

## 5) AgentGuard AI Cyber-Agent Activity Dashboard

**Purpose**
- Detect Mythos-like / vulnerability-research AI agent behaviour.

**Panels**
- Cyber-agent activity over time
- Findings by cyber-agent activity type
- Security tools observed
- Sensitive repository scans
- Exploit-development indicators
- High/critical cyber-agent findings table

**Fields used**
- `ai_sentinel.cyber_agent.activity_type`, `ai_sentinel.cyber_agent.capabilities`
- `ai_sentinel.cyber_agent.security_tools`, `ai_sentinel.cyber_agent.suspicious_keywords`
- `ai_sentinel.cyber_agent.target_paths`, `ai_sentinel.cyber_agent.mitre_tactics`
- `ai_sentinel.cyber_agent.mitre_techniques`, `ai_sentinel.finding.type`
- `host.name`, `process.name`, `ai_sentinel.risk.level`

**Example questions**
- Are AI agents running security tools?
- Are sensitive repositories being scanned?
- Are exploit-development indicators concentrated on specific hosts?

**SOC workflow**
- Review high/critical findings and activity spikes.
- Investigate tooling and target paths used by flagged agents.
- Map behaviour to MITRE tactics/techniques for reporting.

**Tuning suggestions**
- Maintain keyword and tool lists tied to threat intelligence.
- Segment dashboards by engineering vs production endpoints.

## 6) AgentGuard Integration Health Dashboard

**Purpose**
- Monitor pipeline health and telemetry quality.

**Panels**
- Events over time
- Pipeline failures
- Invalid JSON events
- Hosts reporting findings
- Event volume by host
- Findings by `event.action`
- Pipeline failure table

**Fields used**
- `event.module`, `event.dataset`, `tags`
- `ai_sentinel.ingest.error.message`, `error.message`
- `host.name`, `event.action`, `event.outcome`

**Example questions**
- Are any findings failing ingestion/parsing?
- Which hosts are generating parsing failures?
- Are event.action distributions changing unexpectedly?

**SOC workflow**
- Monitor ingest success/failure trends daily.
- Triage failure table and correlate to source hosts.
- Feed recurring parser failure patterns into backlog.

**Tuning suggestions**
- Alert on rising invalid JSON/failure rates.
- Separate parser failures from expected test/noise sources.

## Saved searches

Planned saved searches for operational triage:
- Critical and High AgentGuard Findings
- Untrusted MCP Servers
- Exposed Local LLM Services
- Risky Browser Extensions
- AI Cyber-Agent Activity
- Pipeline Failures
