package util

import (
	"encoding/asn1"
	"encoding/base64"
	"testing"
)

func TestCheckAlgorithmIDParamNotNULL(t *testing.T) {

	testCases := []struct {
		name      string
		checkOID  asn1.ObjectIdentifier
		algorithm string
		errStr    string
	}{
		{
			name:      "valid rsaEncryption",
			checkOID:  OidRSAEncryption,
			algorithm: "MA0GCSqGSIb3DQEBAQUA",
			errStr:    "",
		},
		{
			name:      "valid md2WithRSAEncryption",
			checkOID:  OidMD2WithRSAEncryption,
			algorithm: "MA0GCSqGSIb3DQEBAgUA",
			errStr:    "",
		},
		{
			name:      "valid md5WithRSAEncryption",
			checkOID:  OidMD5WithRSAEncryption,
			algorithm: "MA0GCSqGSIb3DQEBBAUA",
			errStr:    "",
		},
		{
			name:      "valid sha-1WithRSAEncryption",
			checkOID:  OidSHA1WithRSAEncryption,
			algorithm: "MA0GCSqGSIb3DQEBBQUA",
			errStr:    "",
		},
		{
			name:      "valid sha224WithRSAEncryption",
			checkOID:  OidSHA224WithRSAEncryption,
			algorithm: "MA0GCSqGSIb3DQEBDgUA",
			errStr:    "",
		},
		{
			name:      "valid sha256WithRSAEncryption",
			checkOID:  OidSHA256WithRSAEncryption,
			algorithm: "MA0GCSqGSIb3DQEBCwUA",
			errStr:    "",
		},
		{
			name:      "valid sha384WithRSAEncryption",
			checkOID:  OidSHA384WithRSAEncryption,
			algorithm: "MA0GCSqGSIb3DQEBDAUA",
			errStr:    "",
		},
		{
			name:      "valid sha512WithRSAEncryption",
			checkOID:  OidSHA512WithRSAEncryption,
			algorithm: "MA0GCSqGSIb3DQEBDQUA",
			errStr:    "",
		},

		{
			name:      "extra field in algorithm sequence",
			checkOID:  OidRSAEncryption,
			algorithm: "MA8GCSqGSIb3DQEBAQUAAgA=",
			errStr:    "RSA algorithm identifier with trailing data",
		},
		{
			name:      "missing NULL param",
			checkOID:  OidRSAEncryption,
			algorithm: "MAsGCSqGSIb3DQEBAQ==",
			errStr:    "RSA algorithm identifier missing required NULL parameter",
		},
		{
			name:      "NULL param containing data",
			checkOID:  OidRSAEncryption,
			algorithm: "MBQGCSqGSIb3DQEBAQUHTk9UTlVMTA==",
			errStr:    "RSA algorithm identifier with NULL parameter containing data",
		},
		{
			name:      "trailing data after NULL param",
			checkOID:  OidRSAEncryption,
			algorithm: "MBQGCSqGSIb3DQEBAQUATk9UTlVMTA==",
			errStr:    "RSA algorithm identifier with trailing data",
		},
		{
			name:      "context-specific 0 tag in param",
			checkOID:  OidRSAEncryption,
			algorithm: "MA0GCSqGSIb3DQEBAaAA",
			errStr:    "RSA algorithm identifier with non-NULL parameter",
		},
		{
			name:      "wrong algorithm oid",
			algorithm: "MBQGCSqGSIb3DQEBAgUATk9UTlVMTA==",
			errStr:    "error algorithmID to check is not RSA",
		},
		{
			name:      "malformed algorithm sequence",
			checkOID:  OidRSAEncryption,
			algorithm: "MQ0GCSqGSIb3DQEBAQU",
			errStr:    "error reading algorithm",
		},
		{
			name:      "malformed OID",
			checkOID:  OidRSAEncryption,
			algorithm: "MBgTFDEuMi44NDAuMTEzNTQ5LjEuMS4xBQA=",
			errStr:    "error reading algorithm OID",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			algoBytes, _ := base64.StdEncoding.DecodeString(tc.algorithm)

			err := CheckAlgorithmIDParamNotNULL(algoBytes, tc.checkOID)
			if err == nil {
				if tc.errStr != "" {
					t.Errorf("expected error %v was no error", tc.errStr)
				}

				return
			}

			if err.Error() != tc.errStr {
				t.Errorf("expected error %q was %q", tc.errStr, err.Error())
			}
		})
	}
}
