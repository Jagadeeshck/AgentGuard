# Contract validation workflow (v1)

## Canonical contract sources

- Human-readable schema: `contracts/event-schema-v1.md`
- Producer/consumer compatibility rules: `contracts/producer-consumer-contract-v1.md`
- Machine-readable stub schema: `contracts/jsonschema/event-schema-v1.schema.json`

## Canonical fixtures

- Shared fixture set: `examples/sample-findings/v1-sample-events.ndjson`
- Fixture events intentionally remain metadata-only and avoid prompt/body/history payloads.

## How sensor and integrations should use the contract

- `agentguard-sensor` emits contract-aligned `agentguard.*` fields.
- For backward compatibility, sensor events still include `ai_sentinel.*` mirror fields expected by the Elastic package.
- Elastic integration sample events include both namespaces to document current compatibility behavior.

## Validation tests

- `agentguard-sensor/internal/contract/validator_test.go`
  - validates shared fixtures
  - validates sensor-generated event shape
  - rejects prohibited sensitive-content fields

Run with:

```bash
cd agentguard-sensor && go test ./...
```

## Safely adding a new event type

1. Add the new event type in `contracts/event-schema-v1.md` and the JSON schema enum.
2. Add at least one fixture line in `examples/sample-findings/v1-sample-events.ndjson`.
3. Update sensor emission mapping/tests as needed.
4. Ensure Elastic integration sample/pipeline tests still pass with unchanged data stream names.
