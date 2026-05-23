package rules

import "strings"

var secretKeys = []string{"token", "api_key", "apikey", "access_token", "secret", "password", "bearer"}

func RedactSecrets(s string) (string, bool) {
	redacted := s
	found := false
	for _, k := range secretKeys {
		if strings.Contains(strings.ToLower(redacted), k) {
			found = true
			redacted = strings.ReplaceAll(redacted, "=", "=[REDACTED]")
		}
	}
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
