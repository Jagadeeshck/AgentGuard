#!/usr/bin/env bash
set -euo pipefail
launchctl bootout system/com.agentguard.sensor || true
rm -f /Library/LaunchDaemons/com.agentguard.sensor.plist
