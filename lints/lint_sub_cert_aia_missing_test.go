// lint_sub_cert_aia_missing_test.go
package lints

import (
	"testing"
)

func TestSubCertAiaMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertWNoURL.cer"
	desEnum := Error
	out, _ := Lints["sub_cert_aia_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertAiaPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertWBothURL.cer"
	desEnum := Pass
	out, _ := Lints["sub_cert_aia_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
