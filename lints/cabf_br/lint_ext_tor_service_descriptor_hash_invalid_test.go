package lints

import (
	"fmt"
	"testing"
)

func TestTorDescHashInvalid(t *testing.T) {
	testCases := []struct {
		Name            string
		InputFilename   string
		ExpectedResult  LintStatus
		ExpectedDetails string
	}{
		{
			Name:           "Onion subject, no service descriptor extension, before util.CABV201Date",
			InputFilename:  "dnsNameOnionTLD.pem",
			ExpectedResult: NE,
		},
		{
			Name:            "Onion subject, no service descriptor extension, after util.CABV201Date",
			InputFilename:   "onionSANEV.pem",
			ExpectedResult:  Error,
			ExpectedDetails: "certificate contained a .onion domain but is missing a TorServiceDescriptor extension (oid 2.23.140.1.31)",
		},
		{
			Name:            "Onion subject, bad service descriptor, unknown hash algorithm",
			InputFilename:   "onionSANBadServDescUnknownHashAlg.pem",
			ExpectedResult:  Error,
			ExpectedDetails: `TorServiceDescriptor extension (oid 2.23.140.1.31) contained a TorServiceDescriptorHash for Onion URI "https://zmap.onion" with an unknown hash algorithm`,
		},
		{
			Name:            "Onion subject, bad service descriptor, missing hostname",
			InputFilename:   "onionSANBadServDescInvalidUTF8OnionURI.pem",
			ExpectedResult:  Error,
			ExpectedDetails: `TorServiceDescriptor extension (oid 2.23.140.1.31) contained TorServiceDescriptorHash object with Onion URI missing a hostname`,
		},
		{
			Name:            "Onion subject, bad service descriptor, hash alg and hash bit len mismatch",
			InputFilename:   "onionSANBadServDescHashMismatch.pem",
			ExpectedResult:  Error,
			ExpectedDetails: `TorServiceDescriptor extension (oid 2.23.140.1.31) contained a TorServiceDescriptorHash with hash algorithm "SHA256" but only 128 bits of hash not 256`,
		},
		{
			Name:            "Multiple Onion subjects, one missing service descriptor hash entry",
			InputFilename:   "onionSANMissingServDescHash.pem",
			ExpectedResult:  Error,
			ExpectedDetails: `.onion subject domain name "missing.onion" does not have a corresponding TorServiceDescriptorHash for its eTLD+1`,
		},
		{
			Name:            "More service descriptor hash entries than Onion subjects",
			InputFilename:   "onionSANTooManyServDesc.pem",
			ExpectedResult:  Error,
			ExpectedDetails: `TorServiceDescriptor extension (oid 2.23.140.1.31) contained a TorServiceDescriptorHash with a hostname ("other.onion") not present as a subject in the certificate`,
		},
		{
			Name:           "Onion subject, valid service descriptor extension",
			InputFilename:  "onionSANGoodServDesc.pem",
			ExpectedResult: Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			inputPath := fmt.Sprintf("%s%s", testCaseDir, tc.InputFilename)
			result := Lints["e_ext_tor_service_descriptor_hash_invalid"].Execute(ReadCertificate(inputPath))
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
			if result.Details != tc.ExpectedDetails {
				t.Errorf("expected result details %q was %q", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
