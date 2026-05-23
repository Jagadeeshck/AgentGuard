# Watch mode

`agentguard-sensor watch` runs an initial scan and then rescans periodically (default `60s`) to emit NDJSON findings.

## Monitors
- MCP config files
- Browser extension manifests
- Local LLM listening ports
- Startup items
- Suspicious AI-related processes

## Does not monitor
Prompt content, clipboard, browser history, document/file contents, decrypted TLS traffic, secrets, or full environment variables.

## Output and deduplication
- Writes NDJSON append-only to `--output` or to stdout with `--stdout`.
- Stable IDs use SHA-256 with prefixes (`ag-mcp-`, `ag-ext-`, `ag-llm-`, `ag-startup-`, `ag-proc-`).
- Event actions: `finding_detected`, `finding_changed`, `finding_resolved`, `finding_allowed`.
- Statuses: `open`, `changed`, `resolved`, `allowed`.

## State file
Use `--state-file .agentguard-state.json` to persist safe finding metadata (id/type/fingerprint/last_seen/last_risk/last_status) and prevent duplicates across restarts.

## Elastic Agent collection
Configure Elastic Agent filestream input to tail `findings.ndjson` while watch mode appends new lines.

## Service examples
### systemd (Linux user)
```ini
[Unit]
Description=AgentGuard Sensor Watch
After=network.target

[Service]
ExecStart=/usr/local/bin/agentguard-sensor watch --output /var/log/agentguard/findings.ndjson --interval 60s
Restart=on-failure

[Install]
WantedBy=default.target
```

### launchd (macOS)
Write to `/Library/Logs/AgentGuard/findings.ndjson` and run `agentguard-sensor watch --interval 60s`.

### Windows
For MVP run via Task Scheduler or a Windows service wrapper that starts watch mode and rotates logs externally.
