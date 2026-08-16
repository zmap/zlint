package cabf_cs_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCsSubjectOrgRequired(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - Non-EV CS certificate with organizationName in subject",
			InputFilename:  "code_signing/cs_ov_subject_org_present.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - EV CS certificate with organizationName in subject",
			InputFilename:  "code_signing/cs_ev_subject_org_present.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "error - Non-EV CS certificate missing organizationName in subject",
			InputFilename:  "code_signing/cs_ov_subject_org_missing.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "error - EV CS certificate missing organizationName in subject",
			InputFilename:  "code_signing/cs_ev_subject_org_missing.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_cs_requires_org", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
