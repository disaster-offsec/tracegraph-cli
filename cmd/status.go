package cmd

import (
	"flag"
	"fmt"
	"os"

	"tracegraph-cli/internal/client"
	"tracegraph-cli/internal/config"
)

// runStatus handles the `status` command.
func runStatus(args []string) {
	fs := flag.NewFlagSet("status", flag.ContinueOnError)
	backend := fs.String("backend", "", "backend address")

	mustParse(fs, args)

	id := requireOneArg(fs, "usage: tracegraph-cli status <id>")
	url := config.BackendURL(*backend)

	st, err := client.GetStatus(url, id)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Case %s: status=%s\n", st.ID, st.Status)
}
