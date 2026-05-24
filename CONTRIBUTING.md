# Contributing to AgentGuard

Thank you for your interest in contributing to AgentGuard! This document provides guidelines and standards for contributing to the project.

## Code of Conduct

By participating in this project, you agree to abide by our [Code of Conduct](CODE_OF_CONDUCT.md). Please treat everyone with respect and professionalism.

## Copyright and Licensing

AgentGuard is licensed under the **Apache License 2.0**. By submitting a contribution, you agree that your work will be licensed under the same license. Please ensure:

- You have the right to submit your contribution
- Your contribution does not include third-party code incompatible with Apache 2.0
- You add the appropriate copyright header to new files:

```
// Copyright (c) 2025 Jagadeeshck
// SPDX-License-Identifier: Apache-2.0
```

## Getting Started

1. **Fork** the repository on GitHub
2. **Clone** your fork locally:
   ```bash
   git clone https://github.com/<your-username>/AgentGuard.git
   cd AgentGuard
   ```
3. **Create a feature branch** from `main`:
   ```bash
   git checkout -b feature/your-feature-name
   ```
4. Make your changes
5. **Test** your changes thoroughly
6. **Commit** with a clear, descriptive message
7. **Push** to your fork and open a Pull Request

## Development Setup

### Prerequisites

- Go 1.21+
- Docker & Docker Compose
- Elasticsearch 8.x or OpenSearch 2.x (for integration testing)
- Make

### Building

```bash
make build
```

### Running Tests

```bash
make test          # unit tests
make test-integration  # integration tests (requires running Elasticsearch)
```

### Linting

```bash
make lint
```

## How to Contribute

### Reporting Bugs

Use the **Bug Report** issue template. Please include:
- Clear description and steps to reproduce
- Expected vs actual behaviour
- Environment details (OS, Go version, Elasticsearch version)
- Relevant logs

### Suggesting Features

Use the **Feature Request** issue template. Explain the use case and who benefits.

### Asking Questions

Use the **Question / Support Request** issue template, or start a [Discussion](../../discussions).

### Submitting Pull Requests

- One feature or fix per PR — keep changes focused
- Reference the related issue with `Fixes #<issue>` in the PR description
- Fill in the PR template completely
- Ensure all CI checks pass
- Keep commits clean and squash if necessary
- Add or update tests for your changes
- Update documentation if needed

## Commit Message Guidelines

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add support for OpenTelemetry traces
fix: correct Elasticsearch index mapping for agent events
docs: update README with Kubernetes deployment instructions
chore: bump Go version to 1.22
refactor: simplify sensor event batching logic
```

## Code Style

- Follow standard Go conventions (`gofmt`, `golangci-lint`)
- Write clear comments for non-obvious logic
- Keep functions small and focused
- Write tests for new functionality

## Security

If you discover a security vulnerability, **do not** open a public issue. Please follow the [Security Policy](SECURITY.md) for responsible disclosure.

## Sponsoring

If you find AgentGuard valuable, consider [sponsoring the project](../../blob/main/.github/FUNDING.yml) to support ongoing development.

## Questions?

Feel free to open a [Discussion](../../discussions) or a [Question issue](../../issues/new/choose).

---

*AgentGuard is copyright &copy; 2025 Jagadeeshck. Licensed under Apache 2.0.*
