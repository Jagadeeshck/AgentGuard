# Permissions
AgentGuard Sensor reads endpoint metadata for MCP tools, browser extension manifests, local AI service indicators, startup entries, and process metadata.
Some checks may require elevated privileges; when access is denied, the sensor degrades gracefully and emits warning logs.
The sensor does not collect prompt content, clipboard content, browser history, decrypted traffic, private file contents, or secrets.
