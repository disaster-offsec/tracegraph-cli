package cmd

import (
	"flag"
	"fmt"
	"os"
)

// mustParse parses args and exits with code 1 on error.
func mustParse(fs *flag.FlagSet, args []string) {
	if err := fs.Parse(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// requireOneArg returns the single positional argument or exits with code 1.
func requireOneArg(fs *flag.FlagSet, usage string) string {
	if len(fs.Args()) != 1 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	return fs.Args()[0]
}
