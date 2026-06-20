# Elastic Agent sidecar mode for AgentGuard Sensor

Elastic Agent collects AgentGuard Sensor findings through this integration.

## Steps

1. Install AgentGuard Sensor as an OS service.
2. Configure sensor output path:
   - Linux: `/var/log/agentguard/findings.ndjson`
   - macOS: `/Library/Logs/AgentGuard/findings.ndjson`
   - Windows: `C:\ProgramData\AgentGuard\logs\findings.ndjson`
3. Install the AgentGuard integration in Fleet.
4. Set Fleet `paths` to the same sensor output path.
5. Enroll Elastic Agent to the policy.
6. Verify data stream: `logs-agentguard.findings-default`.
7. Verify events include:
   - `event.module: agentguard`
   - `agentguard.finding.type` exists.

## Important wording

In sidecar mode, Elastic Agent **collects** findings files. It does not run the scanner process.
