# Standalone Elastic Agent deployment

Use standalone Elastic Agent mode when a host cannot be managed by Fleet, when configuration is managed by configuration-management tooling, or when you need a minimal deployment that only reads AgentGuard / AI Sentinel NDJSON findings from disk.

This integration package remains only the Elastic package. It does not include the AgentGuard endpoint scanner. The scanner is expected to write one ECS-compatible JSON object per line to a findings file, and Elastic Agent reads that file with the `filestream` input.

## Data stream and pipeline

Use these values unless you intentionally customize the namespace:

- `data_stream.type`: `logs`
- `data_stream.dataset`: `ai_sentinel.findings`
- `data_stream.namespace`: `default`
- Expected data stream: `logs-ai_sentinel.findings-default`
- Package ingest pipeline: `logs-ai_sentinel.findings-default`

The package pipeline is generated from `data_stream/findings/elasticsearch/ingest_pipeline/default.yml`. If your Elastic Stack version or installation process creates a versioned pipeline ID, discover it with one of these methods:

```bash
GET _ingest/pipeline/logs-ai_sentinel.findings*
```

or in Kibana: **Stack Management → Ingest Pipelines**, then search for `logs-ai_sentinel.findings`.

## Complete Linux example

```yaml
outputs:
  default:
    type: elasticsearch
    hosts: ["https://localhost:9200"]
    api_key: "${ELASTIC_API_KEY}"
    ssl.certificate_authorities:
      - /etc/elastic-agent/certs/http_ca.crt

inputs:
  - id: agentguard-ai-sentinel-findings
    type: filestream
    streams:
      - id: agentguard-ai-sentinel-findings-stream
        data_stream:
          type: logs
          dataset: ai_sentinel.findings
          namespace: default
        paths:
          - /var/log/agentguard/findings.ndjson
          - /var/log/ai-sentinel/findings.ndjson
        parsers:
          - ndjson:
              target: ""
              add_error_key: true
              overwrite_keys: true
        pipeline: logs-ai_sentinel.findings-default
        tags:
          - ai_sentinel
          - agentguard
          - ai_activity
          - mcp_monitoring
          - endpoint_visibility
```

## Authentication and TLS

Prefer API key authentication for standalone agents. Create an API key with privileges to write to `logs-ai_sentinel.findings-*` and read package assets required by your operational model, then provide it through an environment variable or your secret-management system:

```bash
export ELASTIC_API_KEY="id:api_key"
```

Configure TLS by pointing `ssl.certificate_authorities` at the Elasticsearch CA certificate. Do not disable certificate verification in production.

## Verify events

1. Confirm Elastic Agent can read the configured findings file.
2. Confirm package assets and the ingest pipeline are installed in Elasticsearch.
3. Query the data stream:

```bash
GET logs-ai_sentinel.findings-default/_search
{
  "size": 5,
  "sort": [{"@timestamp": "desc"}]
}
```

4. In Kibana Discover, select `logs-ai_sentinel.findings-*` and filter for `event.dataset: ai_sentinel.findings`.

## Troubleshooting

- **No events**: verify paths match the scanner output, the file has one JSON object per line, and the Elastic Agent service account can read the file and its parent directories.
- **Pipeline not found**: install or reinstall the package assets, then search ingest pipelines for `logs-ai_sentinel.findings*`.
- **TLS errors**: verify the CA path exists on the host and matches the Elasticsearch HTTP certificate authority.
- **Authentication errors**: check the API key format and privileges.
- **Invalid JSON**: check events tagged `ai_sentinel_invalid_json`; each line must be valid NDJSON.
- **Unexpected raw data**: keep `preserve_original_event` disabled unless debugging in a controlled environment.
