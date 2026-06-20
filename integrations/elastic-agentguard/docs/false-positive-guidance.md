# False Positive Guidance

This package surfaces risky AI-related metadata while avoiding sensitive content collection. Tune on approved business use, trusted tools, host groups, signed paths, extension IDs, and explicit allowlist context. Do not tune by collecting prompts, completions, clipboard data, browsing history, decrypted traffic, secrets, or private file contents.

## Safe examples

| Example | Why it is usually safe | Suggested tuning |
|---|---|---|
| Claude Desktop installed but no dangerous MCP tools. | Installation alone does not imply shell, filesystem, browser, or security-tool capability. | Keep as low risk inventory unless MCP capabilities expand. |
| Ollama bound only to `127.0.0.1`. | Loopback-only local LLM service is not reachable from other hosts. | Allow by process path, loopback destination, owner, and host group. |
| Cursor installed but no shell MCP server. | An approved editor without shell-capable MCP tooling is lower risk. | Allow by signed application path and absence of `shell` / `filesystem` MCP capabilities. |
| Approved internal AI API client. | Business-sanctioned client uses an internal gateway or approved provider. | Set `agentguard.allowed: true` with `allowed_by` and `allowed_at`; scope by process path and destination. |
| Approved security researcher running fuzzing tools. | Defensive research can legitimately use fuzzers or security tools. | Allow by user group, research workstation, approved tools, and ticket or policy reference. |

## Dangerous examples

| Example | Why it is dangerous | Expected response |
|---|---|---|
| Unknown MCP server with shell and filesystem. | Gives an AI client command execution and file access through an untrusted bridge. | Alert high; review server command path, config path, and capabilities. |
| AI agent scans source code and calls external AI API. | Combines repository access with external AI communication and potential sensitive context exposure. | Alert high/critical depending on paths, volume, and approval state. |
| Browser extension with `all_urls` and `nativeMessaging`. | Broad web access plus native host communication can bridge browser and local system access. | Alert high; validate extension ID, deployment policy, and native host registration. |
| Local LLM service bound to `0.0.0.0`. | Service is reachable beyond loopback and may expose local model APIs to other hosts. | Alert high; restrict bind address or firewall exposure. |
| AI agent creates startup persistence. | Persistence can keep an agent or MCP bridge running after reboot/login. | Alert high; validate owner and remove unapproved startup item. |
| AI agent runs security tools from Downloads or temp folder. | Untrusted execution path plus security tooling is a strong behavior signal. | Alert high/critical when paired with shell/filesystem or sensitive paths. |

## High-value review questions

- Is the process, extension, MCP server, or startup item approved for this host and user?
- Does the finding combine multiple risky capabilities, such as MCP plus shell plus filesystem?
- Are target paths limited to a sanctioned project area, or do they include credentials, auth, CI/CD, infrastructure, or runtime control paths?
- Is the network listener loopback-only or reachable from other hosts?
- Did the behavior appear after a configuration change or new startup item?

## Safe tuning practices

- Prefer narrow allowlists over broad rule suppression.
- Keep allowlist metadata in `agentguard.allowed`, `agentguard.allowed_by`, and `agentguard.allowed_at` where possible.
- Use host groups, user groups, signed binaries, extension IDs, MCP server names, approved paths, and approved destinations as tuning dimensions.
- Revisit exceptions periodically, especially for browser extensions and MCP servers.

## What not to tune on

Do not tune by storing or matching prompt contents, completion contents, browser history, clipboard contents, decrypted traffic, secrets, or private file contents. If a false positive requires that level of content to understand, resolve it upstream by improving metadata-only classification.
