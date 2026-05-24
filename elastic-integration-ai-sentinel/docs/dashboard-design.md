# AgentGuard Kibana Dashboard Design (Draft)

This document captures the **design specification** for AgentGuard / AI Sentinel Kibana dashboards before package-compatible saved objects are promoted into `kibana/` assets.

## Scope and constraints

- Data view / index pattern: `logs-ai_sentinel.findings-*`
- No endpoint scanner changes.
- No CI workflow changes.
- No ingest pipeline changes in this milestone.
- Saved objects remain in `dev-assets/kibana_dashboards/` until they are validated as package-compatible.

## Draft assets

Draft Kibana saved-object JSON for six dashboards and six saved searches are stored in:

- `dev-assets/kibana_dashboards/agentguard_ai_activity_overview.dashboard.json`
- `dev-assets/kibana_dashboards/agentguard_mcp_security.dashboard.json`
- `dev-assets/kibana_dashboards/agentguard_local_ai_llm_services.dashboard.json`
- `dev-assets/kibana_dashboards/agentguard_browser_ai_extension_risk.dashboard.json`
- `dev-assets/kibana_dashboards/agentguard_ai_cyber_agent_activity.dashboard.json`
- `dev-assets/kibana_dashboards/agentguard_integration_health.dashboard.json`
- `dev-assets/kibana_dashboards/saved_searches.json`

## Promotion criteria

Draft assets can be moved to package directories (`kibana/dashboard`, `kibana/visualization`, `kibana/search`) only after:

1. Saved-object format and references are validated.
2. `elastic-package lint` passes with the assets in place.
3. `elastic-package build` passes with the assets in place.
4. Any Kibana asset validation command available in the target environment passes.
