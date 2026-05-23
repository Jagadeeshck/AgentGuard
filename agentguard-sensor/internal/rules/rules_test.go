package rules

import "testing"

func TestRedactSecrets(t *testing.T) {
	got, found := RedactSecrets("token=abc123")
	if !found || got == "token=abc123" {
		t.Fatal("expected redaction")
	}
}
func TestRiskScoring(t *testing.T) {
	s, _ := ComputeScore(ScoreInput{MCPShell: true, MCPFilesystem: true})
	if s != 40 {
		t.Fatalf("got %d", s)
	}
}
