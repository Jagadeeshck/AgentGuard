# False Positive Guidance

This package is designed to surface risky AI-related metadata while avoiding sensitive content collection. Analysts should tune on approved business use, trusted tools, and expected environments rather than disabling entire finding families.

## Common benign patterns

| Pattern | Why it can be benign | Suggested tuning |
| --- | --- | --- |
| Approved developer AI tools contacting AI APIs. | Expected use of sanctioned assistants. | Allow by process path, signing identity, provider, and host group. |
| Managed browser extensions with broad permissions. | Enterprise extensions may require broad host permissions. | Allow by extension ID, version, browser profile, and deployment policy. |
| Local LLM bound to loopback only. | Developer or research workflow with local-only exposure. | Lower severity when `destination.ip` is loopback and owner is approved. |
| MCP filesystem access in a sandboxed project directory. | Some coding assistants require project-scoped file access. | Allow by MCP server name, command path, and restricted root directory. |
| Security teams running AI-assisted analysis. | Authorized defensive research. | Allow by user group, workstation tag, ticket reference, and approved tools. |

## High-value review questions

- Is the process, extension, MCP server, or startup item approved for this host and user?
- Does the finding combine multiple risky capabilities, such as MCP plus shell plus filesystem?
- Are target paths limited to a sanctioned project area, or do they include credentials, auth, CI/CD, infrastructure, or runtime control paths?
- Is the network listener loopback-only or reachable from other hosts?
- Did the behavior appear after a configuration change or new startup item?

## Safe tuning practices

- Prefer narrow allowlists over broad rule suppression.
- Keep allowlist metadata in `ai_sentinel.allowed`, `ai_sentinel.allowed_by`, and `ai_sentinel.allowed_at` where possible.
- Use host groups, user groups, signed binaries, extension IDs, MCP server names, and approved paths as tuning dimensions.
- Revisit exceptions periodically, especially for browser extensions and MCP servers.

## What not to tune on

Do not tune by storing or matching prompt contents, completion contents, browser history, clipboard contents, decrypted traffic, secrets, or private file contents. If a false positive requires that level of content to understand, resolve it upstream by improving metadata-only classification.
