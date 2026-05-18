# AI Sentinel Security Rule Reference

The TOML rule drafts were moved to `_dev/security_rules_toml/` because Elastic package security rule assets require package-supported saved-object JSON. This page preserves the rule names and KQL queries for operators until valid package assets are generated.

## AI Agent Writing Exploit-like Files

- Source draft: `_dev/security_rules_toml/ai_agent_exploit_like_files.toml`
- Severity: `high`
- Risk score: `88`
- Description: Detects AI agent findings that indicate exploit-like file creation or modification based on metadata, file names, and behavioural labels.

```kql
event.module: "ai_sentinel" and not ai_sentinel.allowed: true and ai_sentinel.finding.type: "ai_exploit_development_activity" and (file.name: ("*exploit*" or "*poc*" or "*payload*" or "*rop*") or ai_sentinel.cyber_agent.activity_type: "exploit_development" or ai_sentinel.cyber_agent.suspicious_keywords: ("exploit" or "poc" or "payload" or "rop" or "cve"))
```

## AI Agent Sandbox Escape Research Indicators

- Source draft: `_dev/security_rules_toml/ai_agent_sandbox_escape_research.toml`
- Severity: `high`
- Risk score: `86`
- Description: Detects AI agent sandbox escape research indicators from metadata such as activity type, target paths, or behavioural keywords.

```kql
event.module: "ai_sentinel" and not ai_sentinel.allowed: true and ai_sentinel.finding.type: "ai_sandbox_escape_research" and (ai_sentinel.cyber_agent.activity_type: "sandbox_escape_research" or ai_sentinel.cyber_agent.suspicious_keywords: ("sandbox_escape" or "container_escape" or "jailbreak" or "namespace" or "seccomp") or ai_sentinel.cyber_agent.target_paths: ("*/proc/*" or "*/sys/*" or "*/var/run/docker.sock"))
```

## AI Agent Scanning Sensitive Source Code

- Source draft: `_dev/security_rules_toml/ai_agent_sensitive_source_scan.toml`
- Severity: `high`
- Risk score: `78`
- Description: Detects untrusted AI agent findings for sensitive source repository scanning, such as security-critical paths, infra-as-code, auth code, or secrets-related filenames.

```kql
event.module: "ai_sentinel" and not ai_sentinel.allowed: true and ai_sentinel.finding.type: "ai_agent_sensitive_repo_scan" and (ai_sentinel.cyber_agent.target_paths: ("*auth*" or "*security*" or "*crypto*" or "*terraform*" or "*kubernetes*" or "*.env*" or "*secrets*") or ai_sentinel.cyber_agent.suspicious_keywords: ("credential" or "secret" or "token" or "auth" or "crypto"))
```

## AI Agent With Shell and Filesystem MCP Access

- Source draft: `_dev/security_rules_toml/ai_agent_shell_filesystem_mcp.toml`
- Severity: `high`
- Risk score: `85`
- Description: Detects AI agent findings with both shell and filesystem MCP capabilities, which can enable autonomous codebase inspection and tool execution.

```kql
event.module: "ai_sentinel" and not ai_sentinel.allowed: true and ai_sentinel.finding.type: ("ai_agent_shell_tool_use" or "ai_security_tool_mcp_server" or "ai_cyber_agent_activity") and ai_sentinel.cyber_agent.capabilities: "shell" and ai_sentinel.cyber_agent.capabilities: "filesystem"
```

## AI Browser Extension With Broad Permissions

- Source draft: `_dev/security_rules_toml/ai_browser_extension_broad_permissions.toml`
- Severity: `medium`
- Risk score: `55`
- Description: Detects ai browser extension with broad permissions from AI Sentinel findings.

```kql
event.module: "ai_sentinel" and ai_sentinel.finding.type: "browser_extension" and ai_sentinel.extension.permissions: ("<all_urls>" or "scripting" or "webRequest" or "clipboardRead" or "nativeMessaging")
```

## AI Sentinel Critical Finding

- Source draft: `_dev/security_rules_toml/ai_sentinel_critical_finding.toml`
- Severity: `critical`
- Risk score: `90`
- Description: Detects ai sentinel critical finding from AI Sentinel findings.

```kql
event.module: "ai_sentinel" and ai_sentinel.risk.level: "critical"
```

## AI Tool Added to Startup

- Source draft: `_dev/security_rules_toml/ai_tool_added_startup.toml`
- Severity: `medium`
- Risk score: `50`
- Description: Detects ai tool added to startup from AI Sentinel findings.

```kql
event.module: "ai_sentinel" and ai_sentinel.finding.type: "startup_item" and not ai_sentinel.allowed: true
```

## Critical AI Cyber-Agent Activity

- Source draft: `_dev/security_rules_toml/critical_ai_cyber_agent_activity.toml`
- Severity: `critical`
- Risk score: `95`
- Description: Detects critical AI cyber-agent findings across vulnerability research, exploit-development, sandbox escape, shell use, fuzzing, reverse-engineering, and mass codebase analysis behaviours.

```kql
event.module: "ai_sentinel" and ai_sentinel.finding.type: ("ai_cyber_agent_activity" or "ai_vulnerability_research_agent" or "ai_sandbox_escape_research" or "ai_fuzzing_activity" or "ai_reverse_engineering_activity" or "ai_exploit_development_activity" or "ai_security_tool_mcp_server" or "ai_agent_shell_tool_use" or "ai_agent_sensitive_repo_scan" or "ai_agent_mass_codebase_analysis") and (ai_sentinel.risk.level: "critical" or event.severity >= 99 or event.risk_score >= 90)
```

## Local LLM Service Exposed Beyond Localhost

- Source draft: `_dev/security_rules_toml/local_llm_exposed.toml`
- Severity: `high`
- Risk score: `73`
- Description: Detects local llm service exposed beyond localhost from AI Sentinel findings.

```kql
event.module: "ai_sentinel" and ai_sentinel.ai.local_service: true and destination.port: (11434 or 1234 or 7860 or 8000 or 8080) and not destination.ip: ("127.0.0.1" or "::1" or "localhost")
```

## Multiple AI API Connections From Same Process

- Source draft: `_dev/security_rules_toml/multiple_ai_api_connections_threshold.toml`
- Severity: `medium`
- Risk score: `65`
- Description: Detects 10 or more untrusted AI API connection findings from the same host, process, and provider in five minutes.

```kql
event.module: "ai_sentinel" and ai_sentinel.finding.type: "ai_api_connection" and not ai_sentinel.allowed: true
```

## New or Modified MCP Config

- Source draft: `_dev/security_rules_toml/new_modified_mcp_config.toml`
- Severity: `medium`
- Risk score: `47`
- Description: Detects new or modified mcp config from AI Sentinel findings.

```kql
event.module: "ai_sentinel" and event.action: ("mcp_config_created" or "mcp_config_modified")
```

## Possible Mythos-like Vulnerability Research Agent

- Source draft: `_dev/security_rules_toml/possible_mythos_like_vulnerability_research_agent.toml`
- Severity: `high`
- Risk score: `82`
- Description: Detects Mythos-like AI vulnerability research behaviour using multiple behavioural signals. The word mythos is treated only as a weak supporting signal and is never sufficient by itself.

```kql
event.module: "ai_sentinel" and not ai_sentinel.allowed: true and ai_sentinel.finding.type: "ai_vulnerability_research_agent" and (ai_sentinel.cyber_agent.activity_type: ("vulnerability_research" or "exploit_development" or "sandbox_escape_research") or ai_sentinel.cyber_agent.security_tools: * or ai_sentinel.cyber_agent.suspicious_keywords: ("cve" or "vulnerability" or "poc" or "sandbox_escape" or "mythos")) and (ai_sentinel.cyber_agent.capabilities: ("shell" or "filesystem" or "mcp") or ai_sentinel.cyber_agent.target_paths: *)
```

## Unknown Process Connecting to AI API

- Source draft: `_dev/security_rules_toml/unknown_process_ai_api.toml`
- Severity: `medium`
- Risk score: `55`
- Description: Detects unknown process connecting to ai api from AI Sentinel findings.

```kql
event.module: "ai_sentinel" and ai_sentinel.finding.type: "ai_api_connection" and not ai_sentinel.allowed: true and ai_sentinel.ai.provider: *
```

## Untrusted AI Agent Using MCP Browser/Shell Tools

- Source draft: `_dev/security_rules_toml/untrusted_ai_agent_mcp_browser_shell.toml`
- Severity: `high`
- Risk score: `84`
- Description: Detects untrusted AI agent use of MCP browser or shell tooling based on declared capabilities and tool metadata.

```kql
event.module: "ai_sentinel" and not ai_sentinel.allowed: true and ai_sentinel.finding.type: ("ai_security_tool_mcp_server" or "ai_agent_shell_tool_use" or "ai_cyber_agent_activity") and ai_sentinel.cyber_agent.capabilities: "mcp" and ai_sentinel.cyber_agent.capabilities: ("browser" or "shell")
```

## Untrusted AI Agent Running Security Tools

- Source draft: `_dev/security_rules_toml/untrusted_ai_agent_security_tools.toml`
- Severity: `high`
- Risk score: `80`
- Description: Detects untrusted AI agent findings that indicate security tool orchestration or security research capabilities. Behaviour-based and does not rely on an agent name.

```kql
event.module: "ai_sentinel" and not ai_sentinel.allowed: true and ai_sentinel.finding.type: ("ai_security_tool_mcp_server" or "ai_cyber_agent_activity" or "ai_vulnerability_research_agent") and ai_sentinel.cyber_agent.security_tools: *
```

## Untrusted MCP Server With Shell or Filesystem Access

- Source draft: `_dev/security_rules_toml/untrusted_mcp_shell_filesystem.toml`
- Severity: `high`
- Risk score: `75`
- Description: Detects untrusted mcp server with shell or filesystem access from AI Sentinel findings.

```kql
event.module: "ai_sentinel" and ai_sentinel.finding.type: "mcp_server" and not ai_sentinel.allowed: true and ai_sentinel.mcp.capabilities: ("shell" or "filesystem")
```
