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
			Name:           "Onion subject, no service descriptor extension, before util.OnionOnlyEVDate",
			InputFilename:  "dnsNameOnionTLD.pem",
			ExpectedResult: NE,
		},
		{
			Name:            "Onion subject, no service descriptor extension, after util.OnionOnlyEVDate",
			InputFilename:   "onionSANEV.pem",
			ExpectedResult:  Error,
			ExpectedDetails: "certificate contained a .onion domain but is missing a TorServiceDescriptor extension (oid 2.23.140.1.31)",
		},
		{
			Name:            "Onion subject, empty service descriptor extension",
			InputFilename:   "onionSANEmptyServDesc.pem",
			ExpectedResult:  Error,
			ExpectedDetails: "certificate contained an invalid TorServiceDescriptor extension (oid 2.23.140.1.31) - unable to unmarshal outer SEQUENCE",
		},
		{
			Name:            "Onion subject, invalid outer SEQUENCE in service descriptor extension",
			InputFilename:   "onionSANServDescNoOuterSeq.pem",
			ExpectedResult:  Error,
			ExpectedDetails: "certificate contained an invalid TorServiceDescriptor extension (oid 2.23.140.1.31) - invalid outer SEQUENCE",
		},
		{
			Name:            "Onion subject, data trailing outer SEQUENCE in service descriptor extension",
			InputFilename:   "onionSANBadServDescOuterSeqTrailing.pem",
			ExpectedResult:  Error,
			ExpectedDetails: "certificate contained an invalid TorServiceDescriptor extension (oid 2.23.140.1.31) - trailing data after outer SEQUENCE",
		},
		{
			Name:            "Onion subject, data trailing inner TorServiceDescriptorHash SEQUENCE",
			InputFilename:   "onionSANBadServDescInnerSeqTrailing.pem",
			ExpectedResult:  Error,
			ExpectedDetails: "certificate contained an invalid TorServiceDescriptor extension (oid 2.23.140.1.31) - trailing data after TorServiceDescriptorHash",
		},
		{
			Name:            "Onion subject, bad service descriptor onion URI field tag",
			InputFilename:   "onionSANBadServDescOnionURI.pem",
			ExpectedResult:  Error,
			ExpectedDetails: "certificate contained an invalid TorServiceDescriptor extension (oid 2.23.140.1.31) - TorServiceDescriptorHash missing non-compound UTF8String tag",
		},
		{
			Name:            "Onion subject, bad service descriptor onion URI utf8 bytes",
			InputFilename:   "onionSANBadServDescInvalidUTF8OnionURI.pem",
			ExpectedResult:  Error,
			ExpectedDetails: "certificate contained an invalid TorServiceDescriptor extension (oid 2.23.140.1.31) - TorServiceDescriptorHash UTF8String value was not valid UTF-8",
		},
		{
			Name:            "Onion subject, bad service descriptor hash bit length",
			InputFilename:   "onionSANBadServDescBitLen.pem",
			ExpectedResult:  Error,
			ExpectedDetails: "certificate contained an invalid TorServiceDescriptor extension (oid 2.23.140.1.31): invalid TorServiceDescriptorHash subjectPublicKeyHash, bit length is <= 0",
		},
		{
			Name:            "Onion subject, bad service descriptor algorithm field",
			InputFilename:   "onionSANBadServDescAlgorithm.pem",
			ExpectedResult:  Error,
			ExpectedDetails: "certificate contained an invalid TorServiceDescriptor extension (oid 2.23.140.1.31) - error unmarshaling TorServiceDescriptorHash algorithm",
		},
		{
			Name:            "Onion subject, bad service descriptor hash field",
			InputFilename:   "onionSANBadServDescHash.pem",
			ExpectedResult:  Error,
			ExpectedDetails: "certificate contained an invalid TorServiceDescriptor extension (oid 2.23.140.1.31) - error unmarshaling TorServiceDescriptorHash Hash",
		},
		{
			Name:            "Onion subject, bad service descriptor unknown hash algorithm",
			InputFilename:   "onionSANBadServDescUnknownHashAlg.pem",
			ExpectedResult:  Error,
			ExpectedDetails: `certificate contained an invalid TorServiceDescriptor extension (oid 2.23.140.1.31): invalid TorServiceDescriptorHash algorithm "2.16.840.1.101.3.4.2.99"`,
		},
		{
			Name:            "Onion subject, bad service descriptor hash alg and hash bit len mismatch",
			InputFilename:   "onionSANBadServDescHashMismatch.pem",
			ExpectedResult:  Error,
			ExpectedDetails: "certificate contained an invalid TorServiceDescriptor extension (oid 2.23.140.1.31): invalid TorServiceDescriptorHash subjectPublicKeyHash, alg is SHA256 but bit length is 128 not 256",
		},
		{
			Name:            "Multiple Onion subjects, one missing service descriptor hash entry",
			InputFilename:   "onionSANMissingServDescHash.pem",
			ExpectedResult:  Error,
			ExpectedDetails: `.onion domain name "missing.onion" does not have a corresponding TorServiceDescriptorHash`,
		},
		{
			Name:            "More service descriptor hash entries than Onion subjects",
			InputFilename:   "onionSANTooManyServDesc.pem",
			ExpectedResult:  Error,
			ExpectedDetails: "certificate contained more TorServiceDescriptorHash entries than .onion domain names (2 vs 1)",
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
			result := Lints["ext_tor_service_descriptor_hash_invalid"].Execute(ReadCertificate(inputPath))
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
			if result.Details != tc.ExpectedDetails {
				t.Errorf("expected result details %q was %q", tc.ExpectedDetails, result.Details)
			}
		})
	}
}
