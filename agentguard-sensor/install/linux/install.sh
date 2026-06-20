#!/usr/bin/env bash
set -euo pipefail
install -d -m 755 /etc/agentguard /var/log/agentguard /var/lib/agentguard
id -u agentguard >/dev/null 2>&1 || useradd --system --no-create-home --shell /usr/sbin/nologin agentguard || true
if [[ -f /etc/agentguard/config.yml ]]; then
  if ! cmp -s install/linux/config.yml /etc/agentguard/config.yml; then
    backup="/etc/agentguard/config.yml.$(date +%Y%m%d%H%M%S).bak"
    cp /etc/agentguard/config.yml "$backup"
    echo "existing /etc/agentguard/config.yml differs; backed up to $backup and leaving it in place"
  fi
else
  install -m 644 install/linux/config.yml /etc/agentguard/config.yml
fi
install -m 644 install/linux/agentguard-sensor.service /etc/systemd/system/agentguard-sensor.service
chown -R agentguard:agentguard /var/log/agentguard /var/lib/agentguard || true
systemctl daemon-reload
systemctl enable --now agentguard-sensor
