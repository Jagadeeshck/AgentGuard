package process

import "testing"

func TestParsePSOutput(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []Proc
	}{
		{
			name: "parses header and command with spaces",
			in:   "  PID  PPID COMMAND         COMMAND\n 123     1 node            node app.js --token=abc\n",
			want: []Proc{{Name: "node", PID: 123, PPID: 1, Executable: "node", CommandLine: "node app.js --token=abc"}},
		},
		{
			name: "skips malformed and blank lines",
			in:   "PID PPID COMMAND COMMAND\n\nnot-enough\n 42 2 python python -m agent\n",
			want: []Proc{{Name: "python", PID: 42, PPID: 2, Executable: "python", CommandLine: "python -m agent"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParsePSOutput(tt.in)
			if len(got) != len(tt.want) {
				t.Fatalf("len=%d want %d: %#v", len(got), len(tt.want), got)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("got %#v want %#v", got[i], tt.want[i])
				}
			}
		})
	}
}
