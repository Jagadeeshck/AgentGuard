# AI Sentinel Risk Scoring Model

This package accepts producer-supplied `ai_sentinel.risk.score`, `ai_sentinel.risk.level`, and `ai_sentinel.risk.reasons` values and maps them into Elastic Security fields. The future endpoint scanner is responsible for calculating the score before writing NDJSON; this Elastic integration normalizes, validates, and displays the result.

## Risk levels

| Risk level | Score range | Default `event.severity` | Analyst meaning |
|---|---:|---:|---|
| `low` | 0-30 | 21 | Expected or low-impact AI-related metadata. |
| `medium` | 31-60 | 47 | Reviewable behavior with a meaningful but limited risky signal. |
| `high` | 61-85 | 73 | Unapproved behavior with execution, network, file, tool, or exposure risk. |
| `critical` | 86-100 | 99 | High-confidence dangerous combination of multiple strong behavioral signals. |

Scores must be capped at 100 and should not be negative. Producers should keep `ai_sentinel.risk.level` consistent with the score range.

## Additive scoring examples

These are recommended examples for the producer's behavior-based scoring logic:

| Signal | Suggested score contribution |
|---|---:|
| AI provider/API connection | +20 |
| MCP shell capability | +20 |
| MCP filesystem capability | +20 |
| Security tool execution | +15 |
| Large codebase scan | +15 |
| Suspicious cyber keywords in metadata such as filenames, commands, or paths | +15 |
| Exploit-like file writes | +20 |
| Access to browser, email, cloud, or token paths | +20 |
| Process running from temp, downloads, or unknown path | +25 |
| Persistence/startup item | +30 |

## Behavior-based requirements

Scoring must be based on observed behavior and non-sensitive metadata, not on raw user prompts or private content. Strong scores should require combinations such as untrusted process path plus external AI API access, MCP shell plus filesystem capability, broad browser extension permissions plus native messaging, or startup persistence plus agent execution.

A single weak signal must not produce a critical alert. For example, the word `mythos` by itself is not enough to trigger a critical finding. It can only contribute to risk when paired with concrete behavior such as security tool execution, exploit-like file writes, sensitive repository scanning, shell/filesystem MCP access, or public listener exposure.

## Guardrails

- Do not score on prompt content, completion content, clipboard content, browsing history, decrypted traffic, secrets, or private file contents.
- Keep `ai_sentinel.risk.reasons` concise and normalized, for example `mcp_shell_capability`, `external_ai_api`, `temp_path_process`, or `startup_persistence`.
- Treat `ai_sentinel.allowed: true` as triage context. It may reduce detection urgency, but it should not delete the finding.
- Favor transparent, explainable scoring over opaque model-derived labels.
