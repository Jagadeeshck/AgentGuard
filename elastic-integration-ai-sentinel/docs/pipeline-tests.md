# Pipeline Test Fixtures

Pipeline tests live in `data_stream/findings/_dev/test/pipeline/` and include matching `test-<fixture>.log` input fixtures and `test-<fixture>.log-expected.json` expected output files for `elastic-package test pipeline`.

Required MVP coverage:

- `ai_api_connection`
- `mcp_server`
- `browser_extension`
- `startup_item`
- `local_llm_service`
- `invalid_json`
- `redaction`

Additional regression fixtures cover missing optional fields, risk score mapping, and event categorisation.

AI cyber-agent detection pack coverage includes security-tool MCP activity, sensitive repository scanning, exploit-development file metadata, cyber-agent tag derivation, field normalization, and defensive redaction of prohibited cyber-agent payload fields.
