package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/agentguard/agentguard-sensor/internal/config"
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
	for i := 0; i < len(args); i++ {
		if args[i] == "--config" && i+1 < len(args) {
			cp = args[i+1]
			i++
			continue
		}
		if strings.HasPrefix(args[i], "--config=") {
			cp = strings.TrimPrefix(args[i], "--config=")
			continue
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
	f, err := os.Open(*inputPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "validate-output failed:", err)
		os.Exit(1)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	lineNo := 0
	for s.Scan() {
		lineNo++
		line := s.Bytes()
		if len(line) == 0 {
			continue
		}
		var event map[string]any
		if err := json.Unmarshal(line, &event); err != nil {
			fmt.Fprintf(os.Stderr, "invalid JSON at line %d: %v\n", lineNo, err)
			os.Exit(1)
		}
		if err := findings.ValidateContractMap(event); err != nil {
			fmt.Fprintf(os.Stderr, "contract validation failed at line %d: %v\n", lineNo, err)
			os.Exit(1)
		}
	}
	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "validate-output failed:", err)
		os.Exit(1)
	}
	if lineNo == 0 {
		fmt.Fprintln(os.Stderr, "validate-output failed: no events found")
		os.Exit(1)
	}
	fmt.Printf("validated %d event(s) from %s\n", lineNo, *inputPath)
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
	fmt.Println("Usage: agentguard-sensor [--config path] <scan|watch|install-service|uninstall-service|service-status|generate-test-findings|validate-output|version|list-mcp|list-extensions|list-local-ai|list-startup|list-processes>")
}
