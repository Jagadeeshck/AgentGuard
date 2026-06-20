#!/usr/bin/env sh
# Example OTEL environment for a generic AI application emitting metadata-only AgentGuard events.
export OTEL_SERVICE_NAME="example-ai-app"
export OTEL_EXPORTER_OTLP_ENDPOINT="http://agentguard-otel-gateway:4318"
export OTEL_EXPORTER_OTLP_HEADERS="Authorization=Bearer ${AGENTGUARD_OTEL_AUTH_TOKEN}"
export OTEL_RESOURCE_ATTRIBUTES="service.namespace=agentguard,deployment.environment=${AGENTGUARD_ENVIRONMENT:-dev},agentguard.source.type=native_otel,agentguard.source.name=generic_ai_app,event.module=agentguard,event.dataset=agentguard.ai_activity,observer.vendor=AgentGuard,observer.product=AgentGuard,agentguard.privacy.prompt_capture_enabled=false,agentguard.privacy.content_capture_enabled=false,agentguard.privacy.redaction_status=metadata_only"
