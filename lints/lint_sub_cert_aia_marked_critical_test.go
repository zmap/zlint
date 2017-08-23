// lint_sub_cert_aia_missing_test.go
package lints

import (
	"testing"
)

func TestSubCertAiaMarkedCritical(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertAIAMarkedCritical.pem"
	desEnum := Error
	out, _ := Lints["e_sub_cert_aia_marked_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertAiaNotMarkedCritical(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertAIANotMarkedCritical.pem"
	desEnum := Pass
	out, _ := Lints["e_sub_cert_aia_marked_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
