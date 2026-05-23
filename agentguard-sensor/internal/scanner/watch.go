package scanner

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/agentguard/agentguard-sensor/internal/findings"
	"github.com/agentguard/agentguard-sensor/internal/output"
)

type WatchOptions struct {
	OutputPath               string
	Stdout                   bool
	Interval                 time.Duration
	Once                     bool
	EmitResolved             bool
	AllowlistPath, StateFile string
}
type stateRecord struct {
	ID, Type, Fingerprint, LastSeen, LastStatus string
	LastRisk                                    int
}
type watchState struct {
	Findings map[string]stateRecord `json:"findings"`
}

func RunWatch(opts WatchOptions) error {
	if opts.Interval <= 0 {
		opts.Interval = 60 * time.Second
	}
	st := loadState(opts.StateFile)
	run := func() error {
		events, _ := ScanFindings(opts.AllowlistPath)
		now := time.Now().UTC().Format(time.RFC3339)
		cur := map[string]stateRecord{}
		toWrite := []findings.Event{}
		for _, e := range events {
			fp := findings.SafeFingerprint(e)
			prev, ok := st.Findings[e.AIS.Finding.ID]
			cur[e.AIS.Finding.ID] = stateRecord{ID: e.AIS.Finding.ID, Type: e.AIS.Finding.Type, Fingerprint: fp, LastSeen: now, LastRisk: e.AIS.Risk.Score, LastStatus: e.AIS.Finding.Status}
			if !ok {
				toWrite = append(toWrite, e)
				continue
			}
			if prev.Fingerprint != fp {
				e.AIS.Finding.Status = "changed"
				e.Event.Action = "finding_changed"
				toWrite = append(toWrite, e)
			}
		}
		if opts.EmitResolved {
			for id, prev := range st.Findings {
				if _, ok := cur[id]; !ok {
					e := findings.NewEvent(prev.Type, prev.ID, prev.ID, prev.LastRisk, []string{"no longer present"}, map[string]any{"resolved": true})
					e.AIS.Finding.Status = "resolved"
					e.Event.Action = "finding_resolved"
					toWrite = append(toWrite, e)
				}
			}
		}
		st.Findings = cur
		if opts.StateFile != "" {
			_ = saveState(opts.StateFile, st)
		}
		if len(toWrite) > 0 {
			return output.WriteAppend(toWrite, opts.OutputPath, opts.Stdout, false)
		}
		fmt.Fprintln(os.Stderr, "watch: no new or changed findings")
		return nil
	}
	if err := run(); err != nil {
		return err
	}
	if opts.Once {
		return nil
	}
	t := time.NewTicker(opts.Interval)
	defer t.Stop()
	for range t.C {
		if err := run(); err != nil {
			fmt.Fprintln(os.Stderr, "watch cycle failed:", err)
		}
	}
	return nil
}

func loadState(path string) watchState {
	if path == "" {
		return watchState{Findings: map[string]stateRecord{}}
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return watchState{Findings: map[string]stateRecord{}}
	}
	var s watchState
	if json.Unmarshal(b, &s) != nil || s.Findings == nil {
		s.Findings = map[string]stateRecord{}
	}
	return s
}
func saveState(path string, s watchState) error {
	b, _ := json.MarshalIndent(s, "", "  ")
	return os.WriteFile(path, b, 0o600)
}
