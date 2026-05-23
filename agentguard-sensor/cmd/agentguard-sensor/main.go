package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/agentguard/agentguard-sensor/internal/findings"
	"github.com/agentguard/agentguard-sensor/internal/output"
	"github.com/agentguard/agentguard-sensor/internal/scanner"
)

var version = "0.1.0"

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "scan":
		scanCmd(os.Args[2:])
	case "generate-test-findings":
		generateTestFindingsCmd(os.Args[2:])
	case "validate-output":
		validateOutputCmd(os.Args[2:])
	case "version":
		fmt.Printf("agentguard-sensor %s\n", version)
	case "list-mcp", "list-extensions", "list-local-ai", "list-startup", "list-processes":
		listCmd(os.Args[1])
	default:
		usage()
		os.Exit(1)
	}
}

func scanCmd(args []string) {
	fs := flag.NewFlagSet("scan", flag.ExitOnError)
	output := fs.String("output", "", "output NDJSON path")
	stdout := fs.Bool("stdout", false, "write findings to stdout")
	pretty := fs.Bool("pretty", false, "pretty print JSON")
	allowlist := fs.String("allowlist", "", "allowlist YAML path")
	_ = fs.Parse(args)

	opts := scanner.Options{OutputPath: *output, Stdout: *stdout, Pretty: *pretty, AllowlistPath: *allowlist}
	if err := scanner.Run(opts); err != nil {
		fmt.Fprintln(os.Stderr, "scan failed:", err)
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

func listCmd(kind string) { /* unchanged */
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
	fmt.Println("Usage: agentguard-sensor <scan|generate-test-findings|validate-output|version|list-mcp|list-extensions|list-local-ai|list-startup|list-processes>")
}
