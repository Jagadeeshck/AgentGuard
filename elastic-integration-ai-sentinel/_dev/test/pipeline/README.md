Pipeline tests live in `data_stream/findings/test/pipeline` and include matching `test-*.log` input fixtures and `test-*.log-expected.json` expected output files.

Required MVP coverage:

- `test-ai-api-connection.log` / `test-ai-api-connection.log-expected.json`
- `test-mcp-server.log` / `test-mcp-server.log-expected.json`
- `test-browser-extension.log` / `test-browser-extension.log-expected.json`
- `test-startup-item.log` / `test-startup-item.log-expected.json`
- `test-local-llm-service.log` / `test-local-llm-service.log-expected.json`
- `test-invalid-json.log` / `test-invalid-json.log-expected.json`
- `test-redaction.log` / `test-redaction.log-expected.json`

Additional regression fixtures cover missing optional fields, risk score mapping, and event categorisation.

AI cyber-agent detection pack coverage includes security-tool MCP activity, sensitive repository scanning, exploit-development file metadata, cyber-agent tag derivation, field normalization, and defensive redaction of prohibited cyber-agent payload fields.
