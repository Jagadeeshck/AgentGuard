# Safe-vs-Dangerous Scenarios

These scenarios help analysts and package maintainers validate detection logic without collecting sensitive data or building endpoint scanner code.

| Scenario | Safer metadata pattern | Dangerous metadata pattern | Recommended result |
| --- | --- | --- | --- |
| AI API use | Approved process, approved provider, expected host group, no sensitive arguments. | Unknown process contacting AI API from temporary path or unusual user context. | Medium to high depending on trust context. |
| MCP server | Project-scoped filesystem MCP with approved command path. | Untrusted MCP server exposing shell and filesystem tools to an unknown agent. | High or critical when shell and filesystem combine. |
| Browser extension | Managed extension ID and version from enterprise policy. | Extension with `<all_urls>` and native messaging outside approved inventory. | Medium or high based on permissions and allowlist state. |
| Startup item | Approved helper installed by managed software. | New user-writable startup item launching AI automation or MCP bridge. | Medium or high. |
| Local LLM service | Loopback-only listener owned by approved developer workflow. | Non-loopback listener or unexpected service owner. | High or critical when exposed beyond localhost. |
| Cyber-agent research | Authorized security team workstation, ticketed project path, approved tools. | Untrusted agent combining MCP, shell, filesystem, security tools, and sensitive paths. | Critical when multiple dangerous capabilities combine. |
| Sensitive repo scan | Narrow analysis of a project directory with approved business context. | Broad scan of auth, credential, CI/CD, infrastructure, or secret-adjacent paths. | High or critical based on scope and trust. |

## Validation notes

- A safe scenario can still produce a finding; it may be allowlisted or assigned lower risk.
- A dangerous scenario should be represented using metadata-only fields such as paths, tool names, capability labels, and aggregate counts.
- Do not reproduce real exploit logic, prompt text, file contents, secrets, browser history, clipboard data, or decrypted traffic in examples.
