package cmd

import (
	"flag"
	"fmt"

	"tracegraph-cli/internal/config"
)

// runStatus handles the `status` command.
func runStatus(args []string) {
	fs := flag.NewFlagSet("status", flag.ContinueOnError)
	backend := fs.String("backend", "", "backend address")

	mustParse(fs, args)

	id := requireOneArg(fs, "usage: tracegraph-cli status <id>")
	url := config.BackendURL(*backend)
	fmt.Printf("backend=%s, id=%s\n", url, id)
}
