package scanner

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestAllowlistMatch(t *testing.T) {
	if !contains([]string{"a"}, "a") {
		t.Fatal("expected match")
	}
}

func TestLoadAllowlistParsesLocalPorts(t *testing.T) {
	path := filepath.Join(t.TempDir(), "allowlist.yml")
	content := "local_ports:\n  - 11434\n  - '1234'\nprocess_paths: [/usr/local/bin/agent]\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	got := loadAllowlist(path)
	if !reflect.DeepEqual(got.LocalPorts, []int{11434, 1234}) {
		t.Fatalf("unexpected local_ports: %#v", got.LocalPorts)
	}
	if !reflect.DeepEqual(got.ProcessPaths, []string{"/usr/local/bin/agent"}) {
		t.Fatalf("unexpected process_paths: %#v", got.ProcessPaths)
	}
}
