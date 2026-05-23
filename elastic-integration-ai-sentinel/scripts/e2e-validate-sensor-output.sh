#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SAMPLE_FILE="$ROOT_DIR/dev-assets/e2e/agentguard-sensor-findings.ndjson"
ES_URL="${ES_URL:-http://localhost:9200}"
ES_USER="${ES_USER:-elastic}"
ES_PASS="${ES_PASS:-changeme}"

cleanup() {
  elastic-package stack down >/dev/null 2>&1 || true
}
trap cleanup EXIT

command -v elastic-package >/dev/null 2>&1 || { echo "elastic-package is required"; exit 1; }
command -v curl >/dev/null 2>&1 || { echo "curl is required"; exit 1; }
command -v jq >/dev/null 2>&1 || { echo "jq is required"; exit 1; }
[[ -f "$SAMPLE_FILE" ]] || { echo "sample NDJSON not found: $SAMPLE_FILE"; exit 1; }

cd "$ROOT_DIR"
elastic-package build
elastic-package stack up -d --services=elasticsearch
elastic-package install

while IFS= read -r line; do
  [[ -z "$line" ]] && continue
  curl -sS -u "$ES_USER:$ES_PASS" -H 'Content-Type: application/json' \
    -X POST "$ES_URL/logs-ai_sentinel.findings-default/_doc?pipeline=logs-ai_sentinel.findings-default" \
    -d "$line" >/dev/null
done < "$SAMPLE_FILE"

sleep 2

RESP="$(curl -sS -u "$ES_USER:$ES_PASS" "$ES_URL/logs-ai_sentinel.findings-default/_search" -H 'Content-Type: application/json' -d '{"size":20,"query":{"term":{"event.dataset":"ai_sentinel.findings"}}}')"
COUNT="$(echo "$RESP" | jq -r '.hits.total.value')"
[[ "$COUNT" -gt 0 ]] || { echo "no documents found"; exit 1; }

echo "$RESP" | jq -e '.hits.hits[]._source.event.module == "ai_sentinel"' >/dev/null

echo "$RESP" | jq -e '.hits.hits[] | ._source.ai_sentinel.finding.type' >/dev/null

echo "$RESP" | jq -e '.hits.hits[] | ._source.ai_sentinel.risk.score' >/dev/null

for f in 'ai_sentinel.prompt' 'ai_sentinel.prompts' 'ai_sentinel.prompt_content' 'ai_sentinel.completion' 'ai_sentinel.response' 'ai_sentinel.secret' 'ai_sentinel.secrets' 'clipboard' 'browser.history' 'http.request.body.content'; do
  if echo "$RESP" | jq -e --arg path "$f" '
    .hits.hits[] | ._source as $s |
    ($path | split(".")) as $p |
    reduce $p[] as $k ($s; if type=="object" and has($k) then .[$k] else null end) | . != null
  ' >/dev/null; then
    echo "prohibited field present: $f"
    exit 1
  fi
done

echo "E2E validation passed with $COUNT indexed finding(s)."
