# E2E: AgentGuard Sensor -> Elastic integration

## 1) Build package
```bash
cd integrations/elastic-agentguard
elastic-package build
```

## 2) Start Elasticsearch stack
```bash
elastic-package stack up -d --services=elasticsearch
```

## 3) Install package
```bash
elastic-package install
```

## 4) Generate findings from sensor
```bash
agentguard-sensor generate-test-findings --output /tmp/agentguard/findings.ndjson
```

## 5) Ingest findings

### Option A: Elastic Agent filestream + integration
Configure the AgentGuard integration to read `/tmp/agentguard/findings.ndjson` and use pipeline `logs-agentguard.findings-default`.

### Option B: Bulk API quick validation through package pipeline
```bash
cat /tmp/agentguard/findings.ndjson | while IFS= read -r line; do
  curl -s -u elastic:changeme -H 'Content-Type: application/json' \
    -X POST 'http://localhost:9200/logs-agentguard.findings-default/_doc?pipeline=logs-agentguard.findings-default' \
    -d "$line" >/dev/null
 done
```

## 6) Verify in Elasticsearch
```http
GET logs-agentguard.findings-default/_search
```

Expected:
- Events indexed
- `event.module: agentguard`
- `event.dataset: agentguard.findings`
- `agentguard.finding.type` values present
- Redacted values remain redacted
- No prohibited fields are indexed
