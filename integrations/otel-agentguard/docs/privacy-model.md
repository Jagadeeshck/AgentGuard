# Privacy model

AgentGuard is not spyware. It should only be deployed on owned/managed systems or with explicit user and enterprise consent.

Metadata-first is the default. Prompt text, chat content, clipboard content, browser history, decrypted traffic, and raw request/response bodies must not be captured by default.

Defaults:

- `agentguard.privacy.prompt_capture_enabled=false`
- `agentguard.privacy.content_capture_enabled=false`
- `agentguard.privacy.redaction_status=metadata_only`

ChatGPT web monitoring must remain metadata-only unless an enterprise has explicit legal, policy, and user consent. API gateway monitoring should prefer metadata: provider, model, user/team, token counts, status, latency, cost, endpoint, application, and policy decision.

Any optional content capture must be opt-in, policy-bound, redacted, access-controlled, audited, time-limited, and documented.
