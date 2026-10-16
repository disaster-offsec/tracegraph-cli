package cmd

import (
	"flag"
	"fmt"
	"os"

	"tracegraph-cli/internal/bundle"
	"tracegraph-cli/internal/config"
)

// runUpload handles the `upload` command.
func runUpload(args []string) {
	fs := flag.NewFlagSet("upload", flag.ContinueOnError)
	backend := fs.String("backend", "", "backend address")

	mustParse(fs, args)

	path := requireOneArg(fs, "usage: tracegraph-cli upload <path>")
	if err := bundle.Validate(path); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	url := config.BackendURL(*backend)
	fmt.Printf("backend=%s, path=%s\n", url, path)
}
