package output

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/agentguard/agentguard-sensor/internal/findings"
)

func Write(events []findings.Event, path string, toStdout, pretty bool) error {
	if toStdout {
		return writeToStdout(events, pretty)
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
	return writeEvents(f, events, pretty)
}

func WriteAppend(events []findings.Event, path string, toStdout, pretty bool) error {
	if toStdout {
		return writeToStdout(events, pretty)
	}
	if path == "" {
		path = "./findings.ndjson"
	}
	_ = os.MkdirAll(filepath.Dir(path), 0o755)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	return writeEvents(f, events, pretty)
}

func writeToStdout(events []findings.Event, pretty bool) error {
	return writeEvents(os.Stdout, events, pretty)
}
func writeEvents(w interface{ Write([]byte) (int, error) }, events []findings.Event, pretty bool) error {
	for _, e := range events {
		var b []byte
		var err error
		if pretty {
			b, err = json.MarshalIndent(e, "", "  ")
		} else {
			b, err = json.Marshal(e)
		}
		if err != nil {
			return err
		}
		if _, err = w.Write(append(b, '\n')); err != nil {
			return err
		}
	}
	return nil
}
