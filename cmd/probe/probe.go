package probe

import (
	"fmt"
	"os"
)

var filename = fmt.Sprintf("%s/live", os.TempDir())

// Create will remove a file for the liveness check
func Create() error {
	_, err := os.Create(filename)
	return err
}

// Remove will remove the file for the liveness probe
func Remove() error {
	return os.Remove(filename)
}

// Exists checks if the file created for the liveness probe exists
func Exists() bool {
	if _, err := os.Stat(filename); err != nil {
		return false
	}
	return true
}
