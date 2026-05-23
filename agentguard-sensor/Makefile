build:
	go build ./cmd/agentguard-sensor

test:
	go test ./...

run-scan:
	go run ./cmd/agentguard-sensor scan --stdout

run-watch:
	go run ./cmd/agentguard-sensor watch --stdout --once

generate-test-findings:
	go run ./cmd/agentguard-sensor generate-test-findings --output /tmp/agentguard-test.ndjson

validate-output:
	go run ./cmd/agentguard-sensor validate-output --input /tmp/agentguard-test.ndjson
