# Service Installation
## Linux systemd
Run `sudo install/linux/install.sh`, then verify with `systemctl status agentguard-sensor`.
## macOS launchd
Run `sudo install/macos/install.sh`, then verify with `sudo launchctl print system/com.agentguard.sensor`.
## Windows
Run `install/windows/install-service.ps1` in elevated PowerShell, then verify with `Get-Service AgentGuardSensor`.
Findings paths for Elastic Agent collection compatibility:
- `/var/log/agentguard/findings.ndjson`
- `/Library/Logs/AgentGuard/findings.ndjson`
- `C:\ProgramData\AgentGuard\logs\findings.ndjson`
