# Configuration
Use `agentguard-sensor --config /path/to/config.yml watch`.
Default config paths: Linux `/etc/agentguard/config.yml`, macOS `/Library/Application Support/AgentGuard/config.yml`, Windows `C:\ProgramData\AgentGuard\config.yml`.
Privacy guardrails (`collect_prompt_content`, `collect_clipboard`, `collect_browser_history`, `collect_private_file_contents`, `decrypt_traffic`) must remain false in MVP.
