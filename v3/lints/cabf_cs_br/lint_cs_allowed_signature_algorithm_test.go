package cabf_cs_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestCsSignatureAlgorithmNotSupported(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "pass - ECDSA with SHA-256",
			InputFilename:  "code_signing/cs_ecdsa_sha256.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - ECDSA with SHA-384",
			InputFilename:  "code_signing/cs_ecdsa_sha384.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - ECDSA with SHA-512",
			InputFilename:  "code_signing/cs_ecdsa_sha512.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "fail - ECDSA with SHA-1",
			InputFilename:  "code_signing/cs_ecdsa_sha1.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "pass - RSASSA-PKCS1-v1_5 with SHA-256",
			InputFilename:  "code_signing/validCodeSigningCertificate.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - RSASSA-PKCS1-v1_5 with SHA-384",
			InputFilename:  "code_signing/cs_rsa_sha384.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - RSASSA-PKCS1-v1_5 with SHA-512",
			InputFilename:  "code_signing/cs_rsa_sha512.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - RSASSA-PSS with SHA-256",
			InputFilename:  "code_signing/cs_rsa_pss_sha256.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - RSASSA-PSS with SHA-384",
			InputFilename:  "code_signing/cs_rsa_pss_sha384.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "pass - RSASSA-PSS with SHA-512",
			InputFilename:  "code_signing/cs_rsa_pss_sha512.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "fail - SHA-1 with RSA, non-timestamp code signing cert",
			InputFilename:  "code_signing/cs_rsa_sha1_code_signing.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "pass - DSA with SHA-256",
			InputFilename:  "code_signing/cs_dsa_sha256.pem",
			ExpectedResult: lint.Pass,
		},
		{
			Name:           "fail - DSA with SHA-1",
			InputFilename:  "code_signing/cs_dsa_sha1_code_signing.pem",
			ExpectedResult: lint.Error,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_cs_signature_algorithm_not_supported", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v - details: %v", tc.ExpectedResult, result.Status, result.Details)
			}
		})
	}
}
