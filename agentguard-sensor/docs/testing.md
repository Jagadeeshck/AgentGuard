# AgentGuard Sensor Testing

## Unit test coverage areas

- Configuration parsing and Elastic Agent output generation
- Scanner allowlisting and watch-mode dedup/change detection
- Browser extension risk detection
- MCP parsing and capability risk detection
- Rule loading and scoring helpers

## Run tests

```bash
cd agentguard-sensor
go test ./...
```

## Run build

```bash
cd agentguard-sensor
go build ./cmd/agentguard-sensor
```

## Run output validator

Use integration-side validation tooling:

```bash
cd integrations/elastic-agentguard
scripts/e2e-validate-sensor-output.sh
```
