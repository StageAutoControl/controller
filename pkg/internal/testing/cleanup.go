package testing

import (
	"os"
	"testing"
)

// Cleanup a storage path after a test run. Method is meant to be run deferred
func Cleanup(t *testing.T, path string) {
	if err := os.RemoveAll(path); err != nil {
		t.Errorf("failed to remove test storage dir: %v", err)
	}
}
