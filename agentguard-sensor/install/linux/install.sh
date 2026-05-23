#!/usr/bin/env bash
set -euo pipefail
install -d -m 755 /etc/agentguard /var/log/agentguard /var/lib/agentguard
id -u agentguard >/dev/null 2>&1 || useradd --system --no-create-home --shell /usr/sbin/nologin agentguard || true
install -m 644 install/linux/config.yml /etc/agentguard/config.yml
install -m 644 install/linux/agentguard-sensor.service /etc/systemd/system/agentguard-sensor.service
chown -R agentguard:agentguard /var/log/agentguard /var/lib/agentguard || true
systemctl daemon-reload
systemctl enable --now agentguard-sensor
