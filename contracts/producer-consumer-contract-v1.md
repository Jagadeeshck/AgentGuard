# AgentGuard Producer-Consumer Contract v1 (Draft)

## Purpose

Define interoperability expectations between AgentGuard producers (sensor modules) and backend consumer integrations.

## Producer expectations (sensor side)

Producers are expected to:

- emit one JSON object per line (NDJSON)
- emit UTF-8 encoded records
- include required contract fields from the shared schema
- maintain stable event type semantics
- avoid prohibited sensitive-content payload classes by default

## Consumer expectations (backend integrations)

Consumers are expected to:

- accept NDJSON records from configured file/log sources
- parse and validate events with non-fatal handling for malformed lines
- preserve contract field semantics during mapping/enrichment
- isolate backend-specific transforms without mutating event intent

## File-based NDJSON expectations

- line-delimited JSON objects only (no arrays, no multiline single events)
- append-friendly write behavior from producers
- tolerant tail/read behavior from consumers for ongoing files
- clear path configuration ownership per deployment mode

## Schema compatibility rules

- minor/additive schema updates should not break existing consumers
- removal or rename of required fields is a breaking change
- unknown fields must be safely ignored unless explicitly disallowed by backend policy
- contract version must be included for compatibility checks

## Version negotiation and compatibility

- producer emits `agentguard.schema.version`
- consumer declares supported schema versions/ranges in module docs/config
- if schema is unsupported, consumer should:
  1. tag/mark compatibility error,
  2. keep raw event handling bounded and safe,
  3. avoid silent data corruption

## Validation expectations

- producer-side fixtures should cover representative event categories
- consumer-side pipeline/contract tests should validate parse + mapping outcomes
- shared fixtures should be used where possible to reduce drift

## Error handling expectations

- malformed JSON line: do not crash stream processing; emit parse-error metadata where supported
- missing required fields: route to validation-failure path/tag and continue processing
- unsupported version: route to compatibility-failure path/tag and continue processing
- prohibited content fields detected: remove/redact/quarantine per backend policy and log validation signal
