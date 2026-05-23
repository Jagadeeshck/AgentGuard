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
- `ai_sentinel.finding.id`
- `ai_sentinel.finding.type`
- `ai_sentinel.finding.status`
- `ai_sentinel.risk.level`
- `ai_sentinel.risk.score`
- `ai_sentinel.allowed`

Required constants:
- `event.module == ai_sentinel`
- `event.dataset == ai_sentinel.findings`

Prohibited fields:
- `ai_sentinel.prompt`
- `ai_sentinel.prompts`
- `ai_sentinel.prompt_content`
- `ai_sentinel.completion`
- `ai_sentinel.response`
- `ai_sentinel.secret`
- `ai_sentinel.secrets`
- `clipboard`
- `browser.history`
- `http.request.body.content`
