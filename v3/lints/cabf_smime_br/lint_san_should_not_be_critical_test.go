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
			Name:           "pass - certificate with non-critical SAN and non-empty subject",
			InputFilename:  "smime/san_non_critical_non_empty_subject.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "warn - certificate with critical SAN and non-empty subject",
			InputFilename:  "smime/san_critical_non_empty_subject.pem",
			ExpectedResult: lint.Warn,
		},
		{
			Name:           "na - certificate has no SMIME BR policy",
			InputFilename:  "ecdsaP224.pem",
			ExpectedResult: lint.NA,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("w_san_should_not_be_critical", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
