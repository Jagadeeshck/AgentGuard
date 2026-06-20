package rules

import (
	"strings"
	"testing"
)

func TestRedactSecrets(t *testing.T) {
	got, found := RedactSecrets("--token=sk-supersecret123 --config=/etc/app.yml")
	if !found {
		t.Fatal("expected redaction")
	}
	if strings.Contains(got, "sk-supersecret123") {
		t.Fatalf("secret value was not removed: %q", got)
	}
	if !strings.Contains(got, "/etc/app.yml") {
		t.Fatalf("non-secret value should remain untouched: %q", got)
	}
}
func TestRiskScoring(t *testing.T) {
	s, _ := ComputeScore(ScoreInput{MCPShell: true, MCPFilesystem: true})
	if s != 40 {
		t.Fatalf("got %d", s)
	}
}
