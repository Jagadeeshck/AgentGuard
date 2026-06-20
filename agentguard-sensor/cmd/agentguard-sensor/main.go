package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/agentguard/agentguard-sensor/internal/config"
	"github.com/agentguard/agentguard-sensor/internal/contract"
	"github.com/agentguard/agentguard-sensor/internal/findings"
	"github.com/agentguard/agentguard-sensor/internal/output"
	"github.com/agentguard/agentguard-sensor/internal/scanner"
)

var version = "0.1.0"

func main() {
	configPath, args := parseGlobal(os.Args[1:])
	if len(args) < 1 {
		usage()
		os.Exit(1)
	}
	cfg, err := config.Load(configPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "config error:", err)
		os.Exit(1)
	}
	switch args[0] {
	case "scan":
		scanCmd(args[1:], cfg)
	case "watch":
		watchCmd(args[1:], cfg)
	case "install-service", "uninstall-service", "service-status":
		fmt.Println("service command is platform-managed in MVP; use install scripts under install/")
	case "generate-test-findings":
		generateTestFindingsCmd(args[1:])
	case "validate-output":
		validateOutputCmd(args[1:])
	case "config":
		configCmd(args[1:])
	case "doctor":
		doctorCmd(args[1:])
	case "version":
		fmt.Printf("agentguard-sensor %s\n", version)
	case "list-mcp", "list-extensions", "list-local-ai", "list-startup", "list-processes":
		listCmd(args[0])
	default:
		usage()
		os.Exit(1)
	}
}
func parseGlobal(args []string) (string, []string) {
	cp := ""
	out := []string{}
	parsingGlobal := true
	for i := 0; i < len(args); i++ {
		if parsingGlobal && args[i] == "--config" && i+1 < len(args) {
			cp = args[i+1]
			i++
			continue
		}
		if parsingGlobal && strings.HasPrefix(args[i], "--config=") {
			cp = strings.TrimPrefix(args[i], "--config=")
			continue
		}
		if parsingGlobal && !strings.HasPrefix(args[i], "-") {
			parsingGlobal = false
		}
		out = append(out, args[i])
	}
	return cp, out
}

func scanCmd(args []string, cfg config.Config) {
	fs := flag.NewFlagSet("scan", flag.ExitOnError)
	outputPath := fs.String("output", cfg.Output.Path, "output NDJSON path")
	stdout := fs.Bool("stdout", false, "write findings to stdout")
	pretty := fs.Bool("pretty", false, "pretty print JSON")
	allowlist := fs.String("allowlist", cfg.Allowlist.Path, "allowlist YAML path")
	_ = fs.Parse(args)
	opts := scanner.Options{OutputPath: *outputPath, Stdout: *stdout, Pretty: *pretty, AllowlistPath: *allowlist}
	if err := scanner.Run(opts); err != nil {
		fmt.Fprintln(os.Stderr, "scan failed:", err)
		os.Exit(1)
	}
}
func watchCmd(args []string, cfg config.Config) {
	fs := flag.NewFlagSet("watch", flag.ExitOnError)
	outputPath := fs.String("output", cfg.Output.Path, "output NDJSON path")
	stdout := fs.Bool("stdout", false, "write findings to stdout")
	interval := fs.Duration("interval", cfg.Watch.Interval, "rescan interval")
	once := fs.Bool("once", false, "single scan then exit")
	allowlist := fs.String("allowlist", cfg.Allowlist.Path, "allowlist YAML path")
	stateFile := fs.String("state-file", cfg.Watch.StateFile, "optional state file path")
	_ = fs.Parse(args)
	cfg.Output.Path = *outputPath
	cfg.Watch.StateFile = *stateFile
	if err := config.EnsureRuntimePaths(cfg); err != nil && !*stdout {
		fmt.Fprintln(os.Stderr, "watch failed:", err)
		os.Exit(1)
	}
	if err := scanner.RunWatch(scanner.WatchOptions{OutputPath: *outputPath, Stdout: *stdout, Interval: *interval, Once: *once, AllowlistPath: *allowlist, StateFile: *stateFile, EmitResolved: cfg.Watch.EmitResolved}); err != nil {
		fmt.Fprintln(os.Stderr, "watch failed:", err)
		os.Exit(1)
	}
}

func generateTestFindingsCmd(args []string) {
	fs := flag.NewFlagSet("generate-test-findings", flag.ExitOnError)
	outputPath := fs.String("output", "./findings.ndjson", "output NDJSON path")
	_ = fs.Parse(args)
	if err := output.Write(findings.SyntheticTestEvents(), *outputPath, false, false); err != nil {
		fmt.Fprintln(os.Stderr, "generate-test-findings failed:", err)
		os.Exit(1)
	}
	fmt.Printf("wrote synthetic findings to %s\n", *outputPath)
}
func validateOutputCmd(args []string) {
	fs := flag.NewFlagSet("validate-output", flag.ExitOnError)
	inputPath := fs.String("input", "./findings.ndjson", "input NDJSON path")
	_ = fs.Parse(args)
	count, err := contract.ValidateNDJSONFile(*inputPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "validate-output failed:", err)
		os.Exit(1)
	}
	fmt.Printf("validated %d event(s) from %s against AgentGuard finding contract v1\n", count, *inputPath)
}
func listCmd(kind string) {
	results, err := scanner.List(kind)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, r := range results {
		fmt.Println(r)
	}
}
func usage() {
	fmt.Println("Usage: agentguard-sensor [--config path] <scan|watch|config|doctor|install-service|uninstall-service|service-status|generate-test-findings|validate-output|version|list-mcp|list-extensions|list-local-ai|list-startup|list-processes>")
}

func configCmd(args []string) {
	if len(args) == 0 || args[0] != "init" {
		fmt.Fprintln(os.Stderr, "usage: agentguard-sensor config init [--elastic-agent] [--output path]")
		os.Exit(1)
	}
	fs := flag.NewFlagSet("config init", flag.ExitOnError)
	elasticAgent := fs.Bool("elastic-agent", false, "generate an Elastic Agent aligned config")
	outputPath := fs.String("output", "", "output config path")
	_ = fs.Parse(args[1:])
	if !*elasticAgent {
		fmt.Fprintln(os.Stderr, "config init currently requires --elastic-agent")
		os.Exit(1)
	}
	if err := config.WriteElasticAgentConfig(*outputPath); err != nil {
		fmt.Fprintln(os.Stderr, "config init failed:", err)
		os.Exit(1)
	}
	if *outputPath == "" {
		fmt.Printf("wrote Elastic Agent aligned config to %s\n", config.DefaultConfigPath())
		return
	}
	fmt.Printf("wrote Elastic Agent aligned config to %s\n", *outputPath)
}

func doctorCmd(args []string) {
	fs := flag.NewFlagSet("doctor", flag.ExitOnError)
	elasticAgent := fs.Bool("elastic-agent", false, "run Elastic Agent sidecar checks")
	configPath := fs.String("config", "", "sensor config path")
	expectedPath := fs.String("expected-path", "", "expected findings path from Elastic integration")
	_ = fs.Parse(args)
	if !*elasticAgent {
		fmt.Fprintln(os.Stderr, "doctor currently requires --elastic-agent")
		os.Exit(1)
	}
	if *configPath == "" {
		*configPath = config.DefaultConfigPath()
	}
	cfg, err := config.Load(*configPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "doctor failed to load config:", err)
		os.Exit(1)
	}
	if err := cfg.Validate(); err != nil {
		fmt.Fprintln(os.Stderr, "doctor failed:", err)
		os.Exit(1)
	}
	if *expectedPath != "" && cfg.Output.Path != *expectedPath {
		fmt.Fprintf(os.Stderr, "doctor failed: configured output.path %q does not match expected path %q\n", cfg.Output.Path, *expectedPath)
		os.Exit(1)
	}
	if err := os.MkdirAll(filepath.Dir(cfg.Output.Path), 0o755); err != nil {
		fmt.Fprintln(os.Stderr, "doctor failed to ensure output directory:", err)
		os.Exit(1)
	}
	f, err := os.OpenFile(cfg.Output.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "doctor failed: output path not writable:", err)
		os.Exit(1)
	}
	_ = f.Close()
	if stat, err := os.Stat(cfg.Output.Path); err == nil && stat.Size() > 0 {
		validateOutputCmd([]string{"--input", cfg.Output.Path})
	}
	fmt.Printf("doctor checks passed for %s\n", *configPath)
}
