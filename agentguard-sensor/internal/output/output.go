package output

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/agentguard/agentguard-sensor/internal/findings"
)

func Write(events []findings.Event, path string, toStdout, pretty bool) error {
	enc := func(e findings.Event) ([]byte, error) {
		if pretty {
			return json.MarshalIndent(e, "", "  ")
		}
		return json.Marshal(e)
	}
	if toStdout {
		for _, e := range events {
			b, _ := enc(e)
			_, _ = os.Stdout.Write(append(b, '\n'))
		}
		return nil
	}
	if path == "" {
		path = "./findings.ndjson"
	}
	_ = os.MkdirAll(filepath.Dir(path), 0o755)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, e := range events {
		b, _ := enc(e)
		_, _ = f.Write(append(b, '\n'))
	}
	return nil
}
