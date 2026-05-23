#!/usr/bin/env bash
set -euo pipefail
systemctl disable --now agentguard-sensor || true
rm -f /etc/systemd/system/agentguard-sensor.service
systemctl daemon-reload
