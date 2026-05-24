# AgentGuard MVP Backlog (Engineering Epics)

## Epic 1: Product/repo positioning

- **Goal:** Align repository narrative with AgentGuard product architecture.
- **Why it matters:** prevents module confusion and integration-only framing.
- **Tasks:**
  - rewrite root README with product/module split
  - add root architecture/roadmap/modules/privacy docs
  - define root shared contract references
- **Dependencies:** existing module docs inventory.
- **Definition of done:** root docs accurately describe current module boundaries and Elastic-first scope.

## Epic 2: Shared event contract and schema alignment

- **Goal:** establish stable producer/consumer event expectations.
- **Why it matters:** reduces integration drift and ingest breakage.
- **Tasks:**
  - define canonical event schema draft (v1)
  - define producer-consumer compatibility rules
  - map required/recommended fields to ECS alignment guidance
- **Dependencies:** sensor output behavior and Elastic mapping constraints.
- **Definition of done:** contract docs are versioned, referenced by modules, and usable for validation.

## Epic 3: Sensor inventory and metadata coverage

- **Goal:** improve endpoint metadata coverage quality in sensor output.
- **Why it matters:** observability value depends on consistent and useful findings.
- **Tasks:**
  - review existing finding categories and gaps
  - standardize event typing and identifiers
  - add/refresh synthetic examples and tests for major finding classes
- **Dependencies:** Epic 2 contracts.
- **Definition of done:** key finding classes produce consistent fields and pass contract checks.

## Epic 4: Elastic dashboards and detection workflows

- **Goal:** mature Elastic analyst workflows for AgentGuard findings.
- **Why it matters:** ingestion without triage/detection workflow reduces operational value.
- **Tasks:**
  - refine dashboard placeholders toward package-compatible assets
  - improve detection-rule drafts and validation matrix
  - verify ingest pipeline field mappings for dashboard/rule consumption
- **Dependencies:** Epics 2 and 3.
- **Definition of done:** dashboards/rules are documented, testable, and aligned with current event model.

## Epic 5: Browser AI visibility MVP

- **Goal:** deliver browser-related AI activity metadata coverage in a controlled scope.
- **Why it matters:** browser tooling is a high-frequency AI interaction surface.
- **Tasks:**
  - define browser extension and permission finding conventions
  - implement/validate minimal metadata coverage path
  - add sample events and pipeline validation fixtures
- **Dependencies:** Epics 2 and 3; privacy model constraints.
- **Definition of done:** browser metadata findings flow end-to-end without violating privacy model.

## Epic 6: Correlation and policy/risk events

- **Goal:** improve risk interpretation with correlation and policy context.
- **Why it matters:** single findings can be noisy without session/timeline/policy context.
- **Tasks:**
  - introduce correlation identifiers and relationship conventions
  - define policy match/risk event types
  - update docs and validation assets for governance-oriented workflows
- **Dependencies:** Epics 2-5.
- **Definition of done:** correlated and policy-aware events are documented and consumable by backend modules.
