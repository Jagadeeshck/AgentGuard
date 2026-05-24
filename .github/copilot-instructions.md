# Copilot review instructions for AgentGuard

AgentGuard is an AI activity observability platform with:
- `agentguard-sensor/` = endpoint-side producer
- `elastic-integration-ai-sentinel/` = Elastic integration package
- `contracts/` = producer/consumer contracts
- `docs/` = product, privacy, roadmap, and implementation guidance
- `dev-assets/` and `examples/` = fixtures, samples, and validation helpers

When reviewing pull requests, prioritize correctness, security, privacy, and contract stability over style nits.

## What to check first
1. Does the change preserve the metadata-first privacy model?
2. Does it avoid collecting prompts, clipboard contents, browsing history, decrypted traffic, secrets, or credentials unless explicitly documented and approved?
3. Does it preserve redaction and sanitization behavior?
4. Does it keep emitted event shape stable, or clearly document/version any contract changes?
5. Does it keep Elastic package fields, ECS mappings, ingest pipelines, dashboards, and fixtures aligned?

## Sensor review rules
- Flag any new collection of sensitive content.
- Flag unsafe command execution, shell injection, path traversal, unsafe temp-file handling, excessive permissions, and weak input validation.
- Check OS-specific detection logic for false positives, race conditions, and noisy telemetry.
- Ensure findings are structured, deterministic, and testable.
- Prefer minimal collection, clear field naming, and explicit redaction.

## Elastic integration review rules
- Verify ECS alignment and field naming consistency.
- Flag mapping drift between package fields, pipelines, dashboards, sample events, and docs.
- Check ingest pipelines for parsing failures, silent drops, and unhandled malformed events.
- Ensure package assets remain compatible and versioned appropriately.

## Contracts and tests
- Any change to event structure, field names, required fields, or package expectations should update:
  - `contracts/`
  - fixtures/examples
  - tests
  - relevant docs/release notes
- Flag PRs that change behavior without tests or fixtures.

## Security and repo hygiene
- Flag secrets, tokens, credentials, private URLs, customer data, or real sensitive samples.
- Prefer pinned actions, least-privilege workflow permissions, and safe defaults in CI.
- Flag dependency or workflow changes that weaken branch protections, scanning, or release integrity.

## Review style
- Focus on substantive issues, not trivial formatting.
- Explain why an issue matters and suggest the safest fix.
- If a change looks intentional but risky, ask for explicit justification in the PR.
