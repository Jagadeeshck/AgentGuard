# AgentGuard to Elastic Contract v0.1

This contract defines the metadata-only NDJSON events that a future AgentGuard or AI Sentinel producer may write for this Elastic integration. It is intentionally producer-agnostic so the integration can be validated before endpoint scanner code exists.

## Transport

- Format: newline-delimited JSON, one finding per line.
- Destination: file path configured in the Fleet policy for the `ai_sentinel.findings` data stream.
- Encoding: UTF-8.
- Timestamps: ISO 8601 UTC in `@timestamp`.
- Schema namespace: integration-specific fields under `ai_sentinel.*` plus ECS fields such as `host.*`, `process.*`, `file.*`, `destination.*`, `network.*`, `user.*`, and `related.*`.

## Required fields

Every event must include:

- `@timestamp`
- `ai_sentinel.finding.id`
- `ai_sentinel.finding.type`
- `ai_sentinel.risk.level`
- `ai_sentinel.risk.score`
- `ai_sentinel.allowed`

Strongly recommended fields:

- `host.name`
- `ai_sentinel.finding.name`
- `ai_sentinel.finding.status`
- `ai_sentinel.finding.confidence`
- `ai_sentinel.risk.reasons`

## Privacy and data minimization

The producer must emit only metadata needed for detection, triage, and asset correlation. It must not emit:

- Prompt text, completion text, conversation transcripts, or tool call payload content.
- API keys, OAuth tokens, passwords, cookies, session IDs, private keys, or decrypted traffic.
- Clipboard content, browsing history, page contents, form contents, or private file contents.
- Exploit payload source code or step-by-step offensive instructions.

Paths, command names, process names, tool names, model family hints, aggregate counts, and normalized behavior labels are allowed when they do not reveal sensitive content.

## Redaction requirements

Before writing NDJSON, producers must redact secrets from command lines, URLs, file paths, and argument arrays. Use a stable placeholder such as `[REDACTED]` and preserve enough context for detection.

Examples:

- Allowed: `/usr/local/bin/agent --provider openai --api-key [REDACTED]`
- Not allowed: `/usr/local/bin/agent --provider openai --api-key real-token-value`

## Finding lifecycle

Use `ai_sentinel.finding.status` values consistently:

- `open`: current finding requiring attention.
- `acknowledged`: reviewed but still present.
- `closed`: no longer present or accepted through policy.

Use `ai_sentinel.allowed: true` only when the finding matches an approved local allowlist or policy exception. Include `ai_sentinel.allowed_by` and `ai_sentinel.allowed_at` when available.

## Versioning

This document is contract version `v0.1`. Future incompatible changes should add a new contract document rather than changing the meaning of existing fields.
