package cmd

import (
	"flag"
	"fmt"

	"tracegraph-cli/internal/config"
)

// runGet handles the `get` command.
func runGet(args []string) {
	fs := flag.NewFlagSet("get", flag.ContinueOnError)
	backend := fs.String("backend", "", "backend address")

	mustParse(fs, args)

	id := requireOneArg(fs, "usage: tracegraph-cli get <id>")
	url := config.BackendURL(*backend)
	fmt.Printf("backend=%s, id=%s\n", url, id)
}
