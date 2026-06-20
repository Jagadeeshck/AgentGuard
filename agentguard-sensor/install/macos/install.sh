#!/usr/bin/env bash
set -euo pipefail
install -d -m 755 "/Library/Application Support/AgentGuard" "/Library/Logs/AgentGuard"
config_path="/Library/Application Support/AgentGuard/config.yml"
if [[ -f "$config_path" ]]; then
  if ! cmp -s install/macos/config.yml "$config_path"; then
    backup="$config_path.$(date +%Y%m%d%H%M%S).bak"
    cp "$config_path" "$backup"
    echo "existing $config_path differs; backed up to $backup and leaving it in place"
  fi
else
  install -m 644 install/macos/config.yml "$config_path"
fi
install -m 644 install/macos/com.agentguard.sensor.plist /Library/LaunchDaemons/com.agentguard.sensor.plist
launchctl bootstrap system /Library/LaunchDaemons/com.agentguard.sensor.plist
launchctl enable system/com.agentguard.sensor
