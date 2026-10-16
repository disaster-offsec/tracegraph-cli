package bundle

import (
	"fmt"
	"os"
	"path/filepath"
)

// RequiredFiles lists files that must be present in a case bundle.
var RequiredFiles = []string{
	"alert.json",
	"dns.log",
	"conn.log",
	"ground_truth.json",
}

// Validate checks that folder contains a valid case bundle.
func Validate(folder string) error {
	folderInfo, err := os.Stat(folder)
	if err != nil {
		return fmt.Errorf("folder not found: %w", err)
	}
	if !folderInfo.IsDir() {
		return fmt.Errorf("%s is not a directory", folder)
	}

	for _, name := range RequiredFiles {
		path := filepath.Join(folder, name)
		fileInfo, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("missing file: %s", name)
			}
			return fmt.Errorf("cannot stat %s: %w", name, err)
		}
		if fileInfo.IsDir() {
			return fmt.Errorf("%s is a directory, not a file", name)
		}
		if fileInfo.Size() == 0 {
			return fmt.Errorf("%s is empty", name)
		}
	}
	return nil
}
