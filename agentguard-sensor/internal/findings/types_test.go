package findings

import "testing"

func TestValidateEvent(t *testing.T) {
	e := NewEvent("mcp_server", "test", 20, []string{"x"}, nil)
	if err := ValidateRequired(e); err != nil {
		t.Fatal(err)
	}
}
