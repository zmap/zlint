package zlint

import (
	"strings"
	"testing"

	"github.com/zmap/zlint/v3/lint"
)

func TestLintNames(t *testing.T) {
	allowedPrefixes := []string{
		"n_", // lints.Notice
		"w_", // lints.Warn
		"e_", // lints.Error
	}

	for _, name := range lint.GlobalRegistry().Names() {
		var valid bool
		for _, prefix := range allowedPrefixes {
			if strings.HasPrefix(name, prefix) {
				valid = true
				break
			}
		}
		if !valid {
			t.Errorf("lint name %q does not start with an allowed prefix (%v)\n",
				name, allowedPrefixes)
		}
	}
}
