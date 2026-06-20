# AgentGuard Sensor contract test for Elastic ingestion

Before ingestion into Elastic, each NDJSON event must satisfy this contract.

Required fields:
- `@timestamp`
- `ecs.version`
- `event.module`
- `event.dataset`
- `event.kind`
- `event.category`
- `event.type`
- `event.action`
- `event.outcome`
- `event.risk_score`
- `host.name`
- `observer.vendor`
- `observer.product`
- `observer.type`
- `agentguard.finding.id`
- `agentguard.finding.type`
- `agentguard.finding.status`
- `agentguard.risk.level`
- `agentguard.risk.score`
- `agentguard.allowed`

Required constants:
- `event.module == agentguard`
- `event.dataset == agentguard.findings`

Prohibited fields:
- `agentguard.prompt`
- `agentguard.prompts`
- `agentguard.prompt_content`
- `agentguard.completion`
- `agentguard.response`
- `agentguard.secret`
- `agentguard.secrets`
- `clipboard`
- `browser.history`
- `http.request.body.content`
