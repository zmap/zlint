package cabf_cs_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCsEVSubjectOrgMaxLength(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - EV CS certificate with organizationName within 64 characters",
			InputFilename:  "code_signing/cs_ev_subject_org_present.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - EV CS certificate with organizationName exactly 64 characters",
			InputFilename:  "code_signing/cs_ev_subject_org_64_chars.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "error - EV CS certificate with organizationName exceeding 64 characters",
			InputFilename:  "code_signing/cs_ev_subject_org_65_chars.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "NA - EV CS certificate missing organizationName",
			InputFilename:  "code_signing/cs_ev_subject_org_missing.pem",
			ExpectedResult: lint.NA,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_cs_ev_subject_org_max_length", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
