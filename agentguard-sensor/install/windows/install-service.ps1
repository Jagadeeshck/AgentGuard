$serviceName = "AgentGuardSensor"
$exePath = "C:\Program Files\AgentGuard\agentguard-sensor.exe"
$configPath = "C:\ProgramData\AgentGuard\config.yml"
New-Item -ItemType Directory -Force "C:\ProgramData\AgentGuard\logs" | Out-Null
New-Service -Name $serviceName -BinaryPathName "`"$exePath`" --config `"$configPath`" watch" -DisplayName "AgentGuard Sensor" -StartupType Automatic
Start-Service -Name $serviceName
