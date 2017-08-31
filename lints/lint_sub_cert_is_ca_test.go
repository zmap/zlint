// lint_ca_key_cert_sign_not_set_test.go
package lints

import (
	"testing"
)

func TestSubCertIsNotCA(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/subCertIsNotCA.pem"
	desEnum := Pass
	out := Lints["e_sub_cert_not_is_ca"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertIsCA(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/subCertIsCA.pem"
	desEnum := Error
	out := Lints["e_sub_cert_not_is_ca"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
