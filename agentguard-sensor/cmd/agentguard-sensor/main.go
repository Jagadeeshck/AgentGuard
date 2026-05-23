package main

import (
	"flag"
	"fmt"
	"os"

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
	fmt.Println("Usage: agentguard-sensor <scan|version|list-mcp|list-extensions|list-local-ai|list-startup|list-processes>")
}
