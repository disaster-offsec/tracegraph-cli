package cmd

import (
	"fmt"
)

// runStatus handles the `status` command.
func runStatus(args []string) {
	fmt.Println("status called with:", args)
}
