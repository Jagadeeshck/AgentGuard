# Security and privacy notes

This package is intentionally limited to Elastic integration behavior.

- Elastic Agent only reads configured AgentGuard NDJSON files.
- The integration does not inspect processes directly.
- The integration does not inspect browsers directly.
- The integration does not capture packets.
- The integration does not decrypt traffic.
- The integration does not collect prompt content.
- The integration does not collect clipboard content.
- The integration does not collect browsing history.
- The integration does not store secrets.
- Any secret-looking values should already be redacted by the endpoint scanner and are also redacted by the ingest pipeline where supported.
- `event.original` is disabled by default because it can preserve raw input. Enable `preserve_original_event` only for controlled troubleshooting.

The ingest pipeline also removes prohibited raw-content fields if an upstream producer emits them accidentally.
