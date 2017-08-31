// lint_sub_cert_aia_missing_test.go
package lints

import (
	"testing"
)

func TestSubCertAiaMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertWNoURL.pem"
	desEnum := Error
	out := Lints["e_sub_cert_aia_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertAiaPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertWBothURL.pem"
	desEnum := Pass
	out := Lints["e_sub_cert_aia_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
