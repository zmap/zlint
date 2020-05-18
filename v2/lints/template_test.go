package lints

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var (
	// filesChecked is a global counter of the number of files tested by
	// checkForLeftovers.
	filesChecked int
)

// checkForLeftovers checks the given filename (assumed to be a .go src file)
// contains none of the template leftovers. An error is returned if there is
// a problem opening or reading the file, or if any template leftovers are
// found.
func checkForLeftovers(filename string) error {
	// See the `template` file in the root directory of ZLint.
	// None of these strings should appear outside of the template. They indicate
	// the programmer forgot to replace template text.
	leftovers := []string{
		`"Fill this in..."`,
		`"Change this..."`,
		"// Add conditions for application here",
		"// Add actual lint here",
		"Change this to match source TEXT",
	}

	src, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	filesChecked++
	for _, leftover := range leftovers {
		if bytes.Contains(src, []byte(leftover)) {
			return fmt.Errorf(
				"file %q contains template leftover %q",
				filename, leftover)
		}
	}

	return nil
}

// checkFile is a filepath.WalkFunc handler that checks .go files for leftovers.
func checkFile(path string, info os.FileInfo, err error) error {
	// Abort on any incoming errs from filepath.Walk
	if err != nil {
		return err
	}
	// Don't check directories
	if info.IsDir() {
		return nil
	}
	// Only check .go files
	if !strings.HasSuffix(path, ".go") {
		return nil
	}
	// Don't check the template test file, it has the strings we're checking for
	// by design!
	if strings.HasSuffix(path, "template_test.go") {
		return nil
	}

	// Check the path for leftovers
	return checkForLeftovers(path)
}

// TestLeftoverTemplates tests that no .go files under the current directory
// contain leftovers from the new lint template that are intended to be replaced
// by the programmer.
func TestLeftoverTemplates(t *testing.T) {
	if err := filepath.Walk("./", checkFile); err != nil {
		t.Errorf("%v", err)
	}

	// If no files were checked that means something fishy happened. Perhaps the
	// test was moved to a different directory?
	if filesChecked == 0 {
		t.Fatalf("failed to find any files to check while traversing ./")
	}
}
