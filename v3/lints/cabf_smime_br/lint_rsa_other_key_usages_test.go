package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestRSAOtherKeyUsages(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - cert with digitalSignature KU",
			InputFilename:  "smime/rsa_legacy_digital_signature_ku.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "NE - certificate with valid KUs dated before 2020-09-01",
			InputFilename:  "smime/rsa_multipurpose_valid_ku_august_2023.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "NA - cert without KUs",
			InputFilename:  "smime/without_subject_alternative_name.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "NA - cert with KU extension but no KU bits set",
			InputFilename:  "smime/rsa_no_key_usages.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "Error - Certificate with non-zero KUs without digitalSignature or keyEncipherment KUs",
			InputFilename:  "smime/rsa_multipurpose_cert_sign_ku.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_rsa_other_key_usages", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
