package cabf_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestTorDescHashInvalid(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  lint.LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "Onion subject, no service descriptor extension, before util.CABV201Date",
			InputFilename:  "onionSANEVBefore201.pem",
			ExpectedResult: lint.NE,
		},
		{
			Name:            "Onion subject, no service descriptor extension, after util.CABV201Date",
			InputFilename:   "onionSANEV.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: "certificate contained a .onion domain but is missing a TorServiceDescriptor extension (oid 2.23.140.1.31)",
		},
		{
			Name:            "Onion subject, bad service descriptor, unknown hash algorithm",
			InputFilename:   "onionSANBadServDescUnknownHashAlg.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: `TorServiceDescriptor extension (oid 2.23.140.1.31) contained a TorServiceDescriptorHash for Onion URI "https://zmap.onion" with an unknown hash algorithm`,
		},
		{
			Name:            "Onion subject, bad service descriptor, missing hostname",
			InputFilename:   "onionSANBadServDescInvalidUTF8OnionURI.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: `TorServiceDescriptor extension (oid 2.23.140.1.31) contained TorServiceDescriptorHash object with Onion URI missing a hostname`,
		},
		{
			Name:            "Onion subject, bad service descriptor, hash alg and hash bit len mismatch",
			InputFilename:   "onionSANBadServDescHashMismatch.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: `TorServiceDescriptor extension (oid 2.23.140.1.31) contained a TorServiceDescriptorHash with hash algorithm "SHA256" but only 128 bits of hash not 256`,
		},
		{
			Name:            "Multiple Onion subjects, one missing service descriptor hash entry",
			InputFilename:   "onionSANMissingServDescHash.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: `.onion subject domain name "missing.onion" does not have a corresponding TorServiceDescriptorHash for its eTLD+1`,
		},
		{
			Name:            "More service descriptor hash entries than Onion subjects",
			InputFilename:   "onionSANTooManyServDesc.pem",
			ExpectedResult:  lint.Error,
			ExpectedDetails: `TorServiceDescriptor extension (oid 2.23.140.1.31) contained a TorServiceDescriptorHash with a hostname ("other.onion") not present as a subject in the certificate`,
		},
		{
			Name:           "Onion subject, valid service descriptor extension",
			InputFilename:  "onionSANGoodServDesc.pem",
			ExpectedResult: lint.Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_ext_tor_service_descriptor_hash_invalid", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
			if result.Details != tc.ExpectedDetails {
				t.Errorf("expected result details %q was %q", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
