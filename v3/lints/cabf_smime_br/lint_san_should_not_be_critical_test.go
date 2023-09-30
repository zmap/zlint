package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestSubjectAlternativeNameNotCritical(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "warn - san not critical with empty subject",
			InputFilename:  "smime/san_not_critical_with_empty_subject.pem",
			ExpectedResult: lint.Warn,
		},
		{
			Name:           "pass - cert without a CRL distribution point",
			InputFilename:  "smime/san_not_critical_with_subject.pem",
			ExpectedResult: lint.Pass,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_san_should_not_be_critical", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
