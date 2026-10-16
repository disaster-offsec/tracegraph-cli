package cmd

import (
	"flag"
	"fmt"
	"os"

	"tracegraph-cli/internal/bundle"
	"tracegraph-cli/internal/client"
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

	data, err := bundle.Pack(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	url := config.BackendURL(*backend)

	created, err := client.UploadCase(url, data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Case created: id=%s, status=%s\n", created.ID, created.Status)
}
