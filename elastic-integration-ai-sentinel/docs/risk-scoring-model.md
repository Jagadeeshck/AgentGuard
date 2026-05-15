# Risk Scoring Model

The integration accepts producer-provided risk metadata and normalizes it into Elastic Security fields. The ingest pipeline copies `ai_sentinel.risk.score` to `event.risk_score` and maps `ai_sentinel.risk.level` to `event.severity`.

## Risk levels

| Risk level | Score range | Pipeline severity | Interpretation |
| --- | ---: | ---: | --- |
| `low` | 0-39 | 21 | Informational visibility or approved behavior with low-risk metadata. |
| `medium` | 40-69 | 47 | Reviewable behavior with a limited suspicious signal or moderate exposure. |
| `high` | 70-89 | 73 | Unapproved behavior with meaningful execution, network, file, or tool risk. |
| `critical` | 90-100 | 99 | High-confidence dangerous combination such as untrusted agent plus shell/filesystem/security tooling or exposed service risk. |

## Scoring inputs

Recommended additive signals:

- Capability risk: shell, filesystem, browser automation, MCP tool access, startup persistence, non-loopback listeners.
- Trust context: unapproved provider, untrusted path, unsigned or unknown process, absent allowlist.
- Exposure: broad browser permissions, public bind address, sensitive repository paths, high codebase scan volume.
- Tooling: security tools, fuzzers, reverse engineering tools, or exploit-development metadata.
- Confidence: confidence should raise or lower the score only after behavior is established.

## Scoring guardrails

- Do not assign high or critical risk from a model name, provider name, or the word `mythos` alone.
- Do not use prompt content, completion content, browser history, clipboard content, decrypted traffic, secrets, or private file contents in scoring.
- Treat `ai_sentinel.allowed: true` as context for triage and dashboards, not as proof that the event is harmless.
- Preserve `ai_sentinel.risk.reasons` as concise normalized labels so analysts can understand why a score was assigned.
