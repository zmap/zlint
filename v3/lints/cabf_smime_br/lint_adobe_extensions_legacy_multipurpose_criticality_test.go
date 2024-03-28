package cabf_smime_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestAdobeExtensionsLegacyMultipurposeCriticality(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - mailbox legacy cert with non critical adobe time-stamp extension",
			InputFilename:  "smime/mailboxValidatedLegacyWithNonCriticalAdobeTimeStampExtension.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - organization multipurpose cert with non critical adobe archive rev info extension",
			InputFilename:  "smime/organizationValidatedMultipurposeWithNonCriticalAdobeArchRevInfoExtension.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "NA - non-SMIME BR cert",
			InputFilename:  "smime/domainValidatedWithEmailCommonName.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "NA - non-legacy/multipurpose SMIME BR cert",
			InputFilename:  "smime/organizationValidatedStrictWithAdobeTimeStampExtension.pem",
			ExpectedResult: lint.NA,
		},
		{
			Name:           "NE - certificate dated before effective date",
			InputFilename:  "smime/organizationValidatedLegacyWithAdobeTimeStampExtensionMay2023.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:           "Error - sponsor multipurpose certificate with adobe time-stamp extension marked as critical",
			InputFilename:  "smime/sponsorValidatedMultipurposeWithCriticalAdobeTimeStampExtension.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Error - legacy certificate with adobe archive rev info extension marked as critical",
			InputFilename:  "smime/individualValidatedLegacyWithCriticalAdobeArchRevInfoExtension.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_adobe_extensions_legacy_multipurpose_criticality", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
