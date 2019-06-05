package util

import (
	"encoding/base64"
	"testing"
)

func TestCheckAlgorithmIDParamNotNULL(t *testing.T) {

	testCases := []struct {
		name      string
		algorithm string
		errStr    string
	}{
		{
			name:      "valid algorithm",
			algorithm: "MA0GCSqGSIb3DQEBAQUA",
			errStr:    "",
		},
		{
			name:      "extra field in algorithm sequence",
			algorithm: "MA8GCSqGSIb3DQEBAQUAAgA=",
			errStr:    "RSA algorithm identifier with trailing data",
		},
		{
			name:      "missing NULL param",
			algorithm: "MAsGCSqGSIb3DQEBAQ==",
			errStr:    "RSA algorithm identifier missing required NULL parameter",
		},
		{
			name:      "NULL param containing data",
			algorithm: "MBQGCSqGSIb3DQEBAQUHTk9UTlVMTA==",
			errStr:    "RSA algorithm identifier with NULL parameter containing data",
		},
		{
			name:      "trailing data after NULL param",
			algorithm: "MBQGCSqGSIb3DQEBAQUATk9UTlVMTA==",
			errStr:    "RSA algorithm identifier with trailing data",
		},
		{
			name:      "context-specific 0 tag in param",
			algorithm: "MA0GCSqGSIb3DQEBAaAA",
			errStr:    "RSA algorithm identifier with non-NULL parameter",
		},
		{
			name:      "wrong algorithm oid",
			algorithm: "MBQGCSqGSIb3DQEBAgUATk9UTlVMTA==",
			errStr:    "algorithm OID is not rsaEncryption",
		},
		{
			name:      "malformed algorithm sequence",
			algorithm: "MQ0GCSqGSIb3DQEBAQU",
			errStr:    "error reading algorithm",
		},
		{
			name:      "malformed OID",
			algorithm: "MBgTFDEuMi44NDAuMTEzNTQ5LjEuMS4xBQA=",
			errStr:    "error reading algorithm OID",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			algoBytes, _ := base64.StdEncoding.DecodeString(tc.algorithm)

			err := CheckAlgorithmIDParamNotNULL(algoBytes)
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
