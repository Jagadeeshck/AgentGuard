$serviceName = "AgentGuardSensor"
Stop-Service -Name $serviceName -ErrorAction SilentlyContinue
sc.exe delete $serviceName
