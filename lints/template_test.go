package lints

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/zmap/zlint/lint"
)

// TestLeftoverTemplates tests that no .go files for each of the
// lint.LintSources contain leftovers from the new lint template that are
// intended to be replaced by the programmer.
func TestLeftoverTemplates(t *testing.T) {
	// See the `template` file in the root directory of ZLint.
	// None of these strings should appear outside of the template. They indicate
	// the programmer forgot to replace template text.
	leftovers := []string{
		`"Fill this in..."`,
		`"Change this..."`,
		"// Add conditions for application here",
		"// Add actual lint here",
	}

	for _, lintSrc := range lint.LintSources {
		files, err := ioutil.ReadDir(lintSrc.Directory())
		if err != nil {
			t.Fatalf("Failed to read directory %q", lintSrc.Directory())
		}

		for _, f := range files {
			// Skip non-Go files
			if !strings.HasSuffix(f.Name(), ".go") {
				continue
			}

			srcPath := filepath.Join(lintSrc.Directory(), f.Name())
			src, err := ioutil.ReadFile(srcPath)
			if err != nil {
				t.Errorf("Failed to read src file %q: %v",
					f.Name(), err)
				continue
			}

			for _, leftover := range leftovers {
				if bytes.Contains(src, []byte(leftover)) {
					t.Errorf("Lint %q contains template leftover %q",
						f.Name(), leftover)
				}
			}
		}
	}
}
