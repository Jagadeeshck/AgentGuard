Pipeline tests live in `data_stream/findings/test/pipeline` and include matching `.log` input fixtures and `.json` expected output files.

Required MVP coverage:

- `ai_api_connection`
- `mcp_server`
- `browser_extension`
- `startup_item`
- `local_llm_service`
- `invalid_json`
- `redaction`

Additional regression fixtures cover missing optional fields, risk score mapping, and event categorisation.
