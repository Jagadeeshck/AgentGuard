# Local Elastic Package Lab

Use this lab to validate the AI Sentinel Elastic integration before an AgentGuard / AI Sentinel endpoint scanner exists. The lab exercises only the Elastic integration package and synthetic NDJSON findings; it does not include endpoint scanner logic.

## Prerequisites

- Docker or another container runtime supported by `elastic-package`.
- `elastic-package` installed and available on `PATH`.
- A shell with `make`, `curl`, and `jq` for sample ingestion commands.
- This repository checked out locally.

## Install elastic-package

Install the latest release from the upstream Elastic repository:

```bash
curl -fsSL https://api.github.com/repos/elastic/elastic-package/releases/latest \
  | jq -r '.assets[] | select(.name | test("linux_amd64\\.tar\\.gz$")) | .browser_download_url' \
  | xargs curl -fsSL -o /tmp/elastic-package.tar.gz
mkdir -p /tmp/elastic-package
tar -xzf /tmp/elastic-package.tar.gz -C /tmp/elastic-package
sudo install -m 0755 "$(find /tmp/elastic-package -type f -name elastic-package | head -n 1)" /usr/local/bin/elastic-package
elastic-package version
```

macOS users should select the matching Darwin release asset instead of `linux_amd64`.

## Start the local Elastic stack

From the package directory:

```bash
cd elastic-integration-ai-sentinel
elastic-package stack up
```

When the stack is ready, `elastic-package` prints Elasticsearch and Kibana URLs and credentials. Keep these values available for ingestion and Kibana verification.

## Build and install the package

```bash
elastic-package lint
elastic-package build
elastic-package install
```

You can use the Makefile equivalents:

```bash
make lint
make build
make stack-up
```

## Run validation tests

```bash
elastic-package test pipeline
elastic-package test asset
```

Asset tests require a running stack. If no stack is running, run `elastic-package stack up` first.

## Ingest sample NDJSON

The validation corpus lives in `repo-root/dev-assets/test_data/`. Each file contains one or more synthetic JSON lines matching the `ai_sentinel.*` schema.

To ingest one file directly into the findings data stream:

```bash
export ES_URL="https://localhost:9200"
export ES_USER="elastic"
export ES_PASS="<password from elastic-package stack up>"

while IFS= read -r line; do
  curl -k -u "$ES_USER:$ES_PASS" \
    -H 'Content-Type: application/json' \
    -X POST "$ES_URL/logs-ai_sentinel.findings-default/_doc" \
    -d "$line"
done < repo-root/dev-assets/test_data/mcp-server-dangerous.ndjson
```

To ingest every synthetic sample:

```bash
for file in repo-root/dev-assets/test_data/*.ndjson; do
  while IFS= read -r line; do
    curl -k -u "$ES_USER:$ES_PASS" \
      -H 'Content-Type: application/json' \
      -X POST "$ES_URL/logs-ai_sentinel.findings-default/_doc" \
      -d "$line"
  done < "$file"
done
```

## Verify events

Query the data stream:

```bash
curl -k -u "$ES_USER:$ES_PASS" \
  "$ES_URL/logs-ai_sentinel.findings-default/_search?pretty" \
  -H 'Content-Type: application/json' \
  -d '{"size":5,"sort":[{"@timestamp":"desc"}],"query":{"term":{"event.dataset":"ai_sentinel.findings"}}}'
```

In Kibana Discover, select the data view that includes `logs-ai_sentinel.findings-*` and confirm events appear with:

- `event.dataset: ai_sentinel.findings`
- `event.module: ai_sentinel`
- `ai_sentinel.finding.type`
- `ai_sentinel.risk.level`
- `ai_sentinel.risk.score`

## Verify draft detection rules and dashboard placeholders

1. Open Kibana and confirm the AI Sentinel integration package is installed.
2. Draft TOML rules are development references under `repo-root/dev-assets/security_rules_toml/` and summarized in `docs/security-rules.md`; convert them to supported saved-object JSON before installing them as package assets.
3. Use the KQL in `docs/security-rules.md` for manual validation or for rules created outside this package.
4. Ingest the matching `repo-root/dev-assets/test_data/*.ndjson` sample.
5. Confirm expected alerts match `detection-rule-test-matrix.md`.
6. Dashboard placeholder JSON files are kept under `repo-root/dev-assets/kibana_placeholders/` until converted to supported package saved objects.

## Tear down and clean

```bash
elastic-package stack down
rm -rf build .elastic-package
```

Or use:

```bash
make stack-down
make clean
```
