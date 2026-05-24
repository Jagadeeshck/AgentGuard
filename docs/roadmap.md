# AgentGuard Roadmap (Phased)

This roadmap is grounded in current repository contents and active implementation direction.

## Phase 1: Repo and product positioning cleanup

- **Objective:** Establish AgentGuard product-level framing while preserving existing module paths.
- **Deliverables:** root README repositioning, root docs set, root contracts scaffolding.
- **Dependencies:** current module docs and workflow inventory.
- **Exit criteria:** repository clearly communicates producer/consumer split and Elastic-first (not Elastic-only) scope.

## Phase 2: Sensor hardening and finding contract stabilization

- **Objective:** Improve sensor event quality and stabilize event contract boundaries.
- **Deliverables:** schema clarifications, event field consistency, stronger validation fixtures.
- **Dependencies:** sensor implementation backlog and shared contract docs.
- **Exit criteria:** repeatable sensor output quality and documented compatibility expectations.

## Phase 3: Elastic integration and dashboard maturation

- **Objective:** Mature Elastic ingestion, dashboards, and detection workflows.
- **Deliverables:** refined ingest mappings, promoted dashboard assets, stronger detection validation.
- **Dependencies:** stable producer fields and package test coverage.
- **Exit criteria:** reliable analyst workflow for ingest -> investigate -> detect within Elastic.

## Phase 4: Browser AI visibility MVP

- **Objective:** Add practical browser-AI visibility signals without breaking privacy posture.
- **Deliverables:** explicit event types/fields, validation datasets, baseline dashboards/rules.
- **Dependencies:** sensor instrumentation and schema compatibility updates.
- **Exit criteria:** reproducible browser-related metadata findings in end-to-end flow.

## Phase 5: Correlation, risk scoring, and governance events

- **Objective:** Improve context quality via multi-signal correlation and policy/risk semantics.
- **Deliverables:** correlation identifiers, policy match events, risk model hardening.
- **Dependencies:** phase 2 contract stability and phase 3/4 event maturity.
- **Exit criteria:** event streams support governance-aware triage beyond isolated findings.

## Phase 6: Additional backend/export support

- **Objective:** Extend AgentGuard backend portability beyond Elastic.
- **Deliverables:** backend-neutral contract enforcement and at least one additional consumer/export path design.
- **Dependencies:** shared producer-consumer contracts and normalized event categories.
- **Exit criteria:** documented and testable path for non-Elastic backend consumption.
