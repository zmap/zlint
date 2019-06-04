package lints

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestRSAAlgIDNullParams(t *testing.T) {

	testCases := []struct {
		name           string
		filepath       string
		expectedStatus LintStatus
		details        string
	}{
		{
			name:           "pass cert with NULL params",
			filepath:       "rsawithsha1after2016.pem",
			expectedStatus: Pass,
		},
		{
			name:           "error cert with missing NULL params",
			filepath:       "rsaAlgIDNoNULLParams.pem",
			expectedStatus: Error,
			details:        "certificate contains RSA public key algorithm identifier missing required NULL parameter",
		},
		{
			name:           "error cert with non NULL params",
			filepath:       "rsaKeyWithParameters.pem",
			expectedStatus: Error,
			details:        "certificate contains RSA public key algorithm identifier with non-NULL parameter",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputPath := fmt.Sprintf("%s%s", testCaseDir, tc.filepath)
			result := Lints["e_spki_rsa_encryption_parameter_not_null"].Execute(ReadCertificate(inputPath))
			if result.Status != tc.expectedStatus {
				t.Errorf("expected result %v was %v", tc.expectedStatus, result.Status)
			}

			if result.Details != tc.details {
				t.Errorf("expected error details %q was %q", tc.details, result.Details)
			}
		})
	}
}

func TestRSAAlgIDNullParamsSPKI(t *testing.T) {

	testCases := []struct {
		name           string
		spki           string
		expectedStatus LintStatus
		details        string
	}{
		{
			name:           "rsa 1024",
			spki:           "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDIpCF7/X3fdKx6IDmtq2hCCX0BO0zS+LKZdfIIM0rzJ79NIp9nYKKiWT00LLOXFYp/oZjEv0oVQW9Z9B6x4ce6SfV1858Ibz1UAK3IaQIHxhscl3++XKArI1CqCqHtrTcad4SbvM/PmVsSwSn77BSW2Lm+jcdYFvrSKBO8Ibr0LwIDAQAB",
			expectedStatus: Pass,
			details:        "",
		},
		{
			name:           "rsa 2048",
			spki:           "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA/YoWp0tZ2sBGLGPqK3MvqJfKB4Bh2Yr43U2cvpB/Qr6TfYEs6IQOqN89+o5gJF+DRWTWyaaiufc47oU/jfD04wLzCsBM19gqNUA2xQ4GfRmMHNFiKQadBfPVH9ZroqJb799qyzRnXwsgLQVAGUKuam8WDEeMTKl1wQs3m13UVVN3fTyC0SWrQLuJxfvJ13dRrugc+rA7+/xQd9SgACA8EZiJB+IHY6CLgA2+hl5tq5+wiCK6UHAZuHSIqX6HUKCc+jM+sLlNgYSMD+XavypA3scFoWXbjONjUeJjND1c8Mc57dZkPRO2mVMz7jI0T+dHoD4PLrRc9gYfAqsb8sJfIwIDAQAB",
			expectedStatus: Pass,
			details:        "",
		},
		{
			name:           "spki with extra field in algorithm sequence",
			spki:           "MIGkMBIGCSqGSIb3DQEBAQUAAgMBAAEDgY0AMIGJAoGBANgryKYy5GL/TfPQrVmLRae98Ue/CVh7Ir01rpclhpSggMC0H3aRZ0Yx0BCEtyIecCORcsjpbXk6hXeAD8SVFnXFSnFMyGM/o/JjnCpPmvrLwXFuKIUooCceZRyuB9Vbby1D7SuQsYyvJG2u6Rc6BcG/uByuZTsbWMLZrtaqZ4jxAgMBAAE=",
			expectedStatus: Error,
			details:        "certificate contains RSA public key algorithm identifier with trailing data",
		},
		{
			name:           "spki with trailing data after NULL ",
			spki:           "MIGmMBQGCSqGSIb3DQEBAQUHTk9UTlVMTAOBjQAwgYkCgYEA2CvIpjLkYv9N89CtWYtFp73xR78JWHsivTWulyWGlKCAwLQfdpFnRjHQEIS3Ih5wI5FyyOlteTqFd4APxJUWdcVKcUzIYz+j8mOcKk+a+svBcW4ohSigJx5lHK4H1VtvLUPtK5CxjK8kba7pFzoFwb+4HK5lOxtYwtmu1qpniPECAwEAAQ==",
			expectedStatus: Error,
			details:        "certificate contains RSA public key algorithm identifier with NULL parameter containing trailing data",
		},
		{
			name:           "spki with 0 context-specific in params",
			spki:           "MIGfMA0GCSqGSIb3DQEBAaAAA4GNADCBiQKBgQDYK8imMuRi/03z0K1Zi0WnvfFHvwlYeyK9Na6XJYaUoIDAtB92kWdGMdAQhLciHnAjkXLI6W15OoV3gA/ElRZ1xUpxTMhjP6PyY5wqT5r6y8FxbiiFKKAnHmUcrgfVW28tQ+0rkLGMryRtrukXOgXBv7gcrmU7G1jC2a7WqmeI8QIDAQAB",
			expectedStatus: Error,
			details:        "certificate contains RSA public key algorithm identifier with non-NULL parameter",
		},
		{
			name:           "spki bare algo",
			spki:           "MA0GCSqGSIb3DQEBAQUA",
			expectedStatus: Fatal,
			details:        "error reading pkixPublicKey algorithm",
		},
		{
			name:           "wrong algorithm oid",
			spki:           "MIGfMA0GCSqGSIb3DQEBAgUAA4GNADCBiQKBgQDYK8imMuRi/03z0K1Zi0WnvfFHvwlYeyK9Na6XJYaUoIDAtB92kWdGMdAQhLciHnAjkXLI6W15OoV3gA/ElRZ1xUpxTMhjP6PyY5wqT5r6y8FxbiiFKKAnHmUcrgfVW28tQ+0rkLGMryRtrukXOgXBv7gcrmU7G1jC2a7WqmeI8QIDAQAB",
			expectedStatus: Error,
			details:        "certificate pkixPublicKey algorithm OID is not rsaEncryption",
		},
	}

	inputPath := fmt.Sprintf("%s%s", testCaseDir, "rsawithsha1after2016.pem")
	cert := ReadCertificate(inputPath)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			spki := make([]byte, len([]byte(tc.spki)))
			base64.StdEncoding.Decode(spki, []byte(tc.spki))
			cert.RawSubjectPublicKeyInfo = spki

			result := Lints["e_spki_rsa_encryption_parameter_not_null"].Execute(cert)
			if result.Status != tc.expectedStatus {
				t.Errorf("expected result %v was %v", tc.expectedStatus, result.Status)
			}

			if result.Details != tc.details {
				t.Errorf("expected error details %q was %q", tc.details, result.Details)
			}
		})
	}
}
