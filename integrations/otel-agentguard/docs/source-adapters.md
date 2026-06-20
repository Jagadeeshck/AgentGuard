# Source adapters

- `native_otel`: tools and frameworks emit OTLP directly with `agentguard.*` attributes.
- `api_gateway`: OpenAI API, Azure OpenAI, Anthropic, Gemini, Bedrock, and internal gateways provide metadata from HTTP or app telemetry.
- `browser_metadata`: ChatGPT, Gemini, Claude web, Perplexity, and Copilot web provide approved browser/session metadata only; no prompt or page content by default.
- `mcp_runtime`: MCP server connection and tool invocation metadata, tool category, target path classification, and risk score.
- `local_llm`: endpoint or runtime discovery for Ollama, LM Studio, vLLM, llama.cpp, or similar local services.
- `endpoint_sensor`: existing `agentguard-sensor` output can be converted to OTEL or correlated downstream.
- `saas_audit`: placeholder for vendor-provided enterprise audit exports; no undocumented scraping.

Safe OpenAI patterns are `openai_api_gateway`, `chatgpt_browser_metadata`, and `chatgpt_enterprise_audit_placeholder`.
