# AgentGuard Privacy Model (Current Repository Scope)

## Privacy posture

AgentGuard currently follows a metadata-first collection model designed for enterprise-safe observability.

## Default collection constraints

By default, AgentGuard design and repository guidance enforce:

- no prompt capture
- no clipboard capture
- no browsing history collection
- no decrypted traffic inspection
- no intentional secret/credential capture

## Operational meaning

- Collected events should represent behavioral/technical metadata (for example process, service, config, extension, path, and risk context) rather than user content.
- Backend consumers should process structured findings and avoid introducing raw-content expansion that violates upstream privacy assumptions.
- Debug/forensic controls that preserve raw input must be explicit, controlled, and minimized.

## Enterprise control expectations

- Data minimization is the default baseline.
- Opt-in controls must be explicit and documented before enabling broader collection classes.
- Security and privacy review should accompany any expansion to event payload scope.
- Module docs and contracts should remain consistent with this model.

## Future content-capture rule

Any future content-level capture capability must be:

1. explicitly designed,
2. explicitly documented,
3. explicitly opt-in,
4. bounded by policy and governance controls.

Until those conditions are met, content capture remains out of current scope.
