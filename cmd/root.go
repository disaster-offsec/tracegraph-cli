package cmd

import (
	"fmt"
	"os"
)

// Execute parses CLI arguments and dispatches to the right command.
func Execute() {
	args := os.Args[1:]

	if len(args) == 0 {
		printUsage()
		os.Exit(0)
	}

	cmd := args[0]
	rest := args[1:]

	switch cmd {
	case "upload":
		runUpload(rest)
	case "status":
		runStatus(rest)
	case "get":
		runGet(rest)
	case "--help", "-h", "help":
		printUsage()
		os.Exit(0)
	default:
		fmt.Printf(os.Stderr, "unknown command: %s\n\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

// printUsage prints the CLI usage and available commands.
func printUsage() {
	fmt.Println("Usage: tracegraph-cli <command> [args] [flags]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  upload <path>   upload a case bundle from a folder")
	fmt.Println("  status <id>     show case status")
	fmt.Println("  get <id>        show full case")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  --backend <url>  backend address (default: TRACEGRAPH_BACKEND env or http://localhost:8000)")
	fmt.Println("  --help, -h       show this help")
}
