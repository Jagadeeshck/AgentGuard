# E2E: Elastic Agent + AgentGuard Sensor sidecar

Flow:

AgentGuard Sensor service  
↓ writes findings.ndjson  
Elastic Agent (filestream)  
↓ ingest pipeline  
`logs-ai_sentinel.findings-default`

## Linux

```bash
sudo systemctl enable --now agentguard-sensor
sudo agentguard-sensor config init --elastic-agent --output /etc/agentguard/config.yml
sudo elastic-agent enroll <fleet-url> <enrollment-token>
```

## macOS

```bash
sudo launchctl load /Library/LaunchDaemons/com.agentguard.sensor.plist
sudo agentguard-sensor config init --elastic-agent --output "/Library/Application Support/AgentGuard/config.yml"
sudo elastic-agent enroll <fleet-url> <enrollment-token>
```

## Windows (PowerShell)

```powershell
Start-Service AgentGuardSensor
agentguard-sensor.exe config init --elastic-agent --output "C:\ProgramData\AgentGuard\config.yml"
& "C:\Program Files\Elastic\Agent\elastic-agent.exe" enroll <fleet-url> <enrollment-token>
```

## Validation

- Ensure Fleet `paths` matches sensor `output.path`.
- Query `logs-ai_sentinel.findings-default`.
- Confirm `event.module == "ai_sentinel"` and `ai_sentinel.finding.type` is populated.
