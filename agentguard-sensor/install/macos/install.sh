#!/usr/bin/env bash
set -euo pipefail
install -d -m 755 "/Library/Application Support/AgentGuard" "/Library/Logs/AgentGuard"
install -m 644 install/macos/config.yml "/Library/Application Support/AgentGuard/config.yml"
install -m 644 install/macos/com.agentguard.sensor.plist /Library/LaunchDaemons/com.agentguard.sensor.plist
launchctl bootstrap system /Library/LaunchDaemons/com.agentguard.sensor.plist
launchctl enable system/com.agentguard.sensor
