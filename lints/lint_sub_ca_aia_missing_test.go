// lint_sub_ca_aia_missing_test.go
package lints

import (
	"testing"
)

func TestSubCaAiaMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAAIAMissing.pem"
	desEnum := Error
	out, _ := Lints["e_sub_ca_aia_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCaAiaPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAAIAValid.pem"
	desEnum := Pass
	out, _ := Lints["e_sub_ca_aia_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
