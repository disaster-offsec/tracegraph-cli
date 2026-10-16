package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"tracegraph-cli/internal/client"
	"tracegraph-cli/internal/config"
)

// runGet handles the `get` command.
func runGet(args []string) {
	fs := flag.NewFlagSet("get", flag.ContinueOnError)
	backend := fs.String("backend", "", "backend address")

	mustParse(fs, args)

	id := requireOneArg(fs, "usage: tracegraph-cli get <id>")
	url := config.BackendURL(*backend)

	c, err := client.GetCase(url, id)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	out, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}
