# Field mapping

AgentGuard events SHOULD be emitted as OTEL log attributes. Metrics and traces MAY reuse the same attributes for correlation.

| AgentGuard field | Type | Notes |
| --- | --- | --- |
| `agentguard.cost.estimated_usd` | float | Estimated cost, not billing authority. |
| `agentguard.tokens.input` | long | Input token count when available. |
| `agentguard.tokens.output` | long | Output token count when available. |
| `agentguard.tokens.total` | long | Total token count when available. |
| `agentguard.duration.ms` | long | Request, session, or tool duration. |
| `agentguard.risk.score` | integer | 0-100 risk score. |
| `agentguard.privacy.prompt_capture_enabled` | boolean | Must default to false. |
| `agentguard.privacy.content_capture_enabled` | boolean | Must default to false. |
| `agentguard.attributes` | flattened/object | Flexible source metadata. |
| `agentguard.tool.parameters` | flattened/object | Sanitized tool parameters only. |
| `agentguard.source.raw_metadata` | flattened/object | Metadata only; no raw content. |

Elastic, OpenSearch, Splunk, Datadog, file, and OTLP exporters should preserve `agentguard.*` names whenever practical.
