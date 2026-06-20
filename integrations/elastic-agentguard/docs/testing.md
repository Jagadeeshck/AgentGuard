# Testing Guide

## Pipeline tests structure

Pipeline fixtures live in `integrations/elastic-agentguard/data_stream/findings/_dev/test/pipeline/`.

Each test uses:

- `test-<name>.log`
- `test-<name>.log-expected.json`

Use hyphenated names only (no underscores).

## Generate expected fixtures

```bash
cd integrations/elastic-agentguard
elastic-package stack up -d --services=elasticsearch
elastic-package test pipeline --generate
elastic-package stack down
```

After generation, manually review expected output before commit.

## Safety review checklist

- Fixtures are synthetic only.
- No real personal data.
- No real secrets/tokens/keys.
- No prompt/completion content.
- No clipboard or browser history contents.

## Prohibited fields checklist

Expected outputs for valid events should not include prohibited raw content fields such as:

- `agentguard.prompt`
- `agentguard.prompt_content`
- `agentguard.completion`
- `agentguard.response`
- `agentguard.secret`
- `agentguard.secrets`
- `clipboard`
- `browser.history`
- `http.request.body.content`

## Redaction checklist

Confirm redaction occurs for sensitive values in:

- command-line flags (`--token`, `--api-key`)
- query parameters (`token=`, `api_key=`, `access_token=`)
- bearer tokens
- MCP/server argument lists
- AI endpoint URLs carrying secrets


Canonical naming: `agentguard.*` is the only supported custom namespace for current contracts.


Package contract baseline for tests:
- Package name: `agentguard`
- Dataset: `agentguard.findings`
- Data stream: `logs-agentguard.findings-default`
- Canonical custom namespace: `agentguard.*`
