package lints

import (
	"testing"

	"github.com/zmap/zlint/v3/integration/lints/lint"
)

func TestRegisterLintDeprecated_Lint(t *testing.T) {

	data := []struct {
		inputFile  string
		expectPass bool
	}{
		{inputFile: "testdata/lint_usesRegisterLint.go", expectPass: false},
		{inputFile: "testdata/lint_usesRegisterCertificateLint.go", expectPass: true},
		{inputFile: "testdata/lint_usesRegisterProfile.go", expectPass: true},
		{inputFile: "testdata/lint_usesRegisterRevocationListLint.go", expectPass: true},
	}
	l := &RegisterLintDeprecated{}
	for _, test := range data {
		file := test.inputFile
		want := test.expectPass
		t.Run(file, func(t *testing.T) {
			r, err := lint.RunLintForFile(file, l)
			if err != nil {
				t.Fatal(err)
			}
			if want && r != nil {
				t.Errorf("got unexepcted error result, %s", r)
			} else if !want && r == nil {
				t.Errorf("expected failure but got nothing")
			}
		})
	}
}
