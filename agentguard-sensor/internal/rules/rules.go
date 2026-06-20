package rules

import "regexp"

var secretPattern = regexp.MustCompile(`(?i)(^|\s)([^=\s]*(?:token|api_key|apikey|access_token|secret|password|bearer))=([^\s]*)`)

func RedactSecrets(s string) (string, bool) {
	found := false
	redacted := secretPattern.ReplaceAllStringFunc(s, func(match string) string {
		found = true
		parts := secretPattern.FindStringSubmatch(match)
		return parts[1] + parts[2] + "=[REDACTED]"
	})
	return redacted, found
}

type ScoreInput struct {
	MCPShell, MCPFilesystem, BrowserRisky, LocalExposed, Startup bool
	UnknownRuntimeAI, SecurityKeywords, RedactedSecret, TempPath bool
}

func ComputeScore(i ScoreInput) (int, []string) {
	score := 0
	reasons := []string{}
	add := func(cond bool, points int, reason string) {
		if cond {
			score += points
			reasons = append(reasons, reason)
		}
	}
	add(i.MCPShell, 20, "mcp shell capability")
	add(i.MCPFilesystem, 20, "mcp filesystem capability")
	add(i.BrowserRisky, 15, "risky browser permissions")
	add(i.LocalExposed, 20, "non-loopback local ai service")
	add(i.Startup, 15, "startup persistence")
	add(i.UnknownRuntimeAI, 10, "runtime with ai keywords")
	add(i.SecurityKeywords, 15, "security/cyber keywords")
	add(i.RedactedSecret, 20, "token-like value redacted")
	add(i.TempPath, 10, "execution from temp/downloads path")
	if score > 100 {
		score = 100
	}
	return score, reasons
}
