// lint_cert_extensions_verson_not_3_test.go
package lints

import (
	"testing"
)

func TestExtsV2(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/certVersion2WithExtension.pem"
	desEnum := Error
	out, _ := Lints["e_cert_extensions_version_not_3"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestExtsV3(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caBasicConstCrit.pem"
	desEnum := Pass
	out, _ := Lints["e_cert_extensions_version_not_3"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestNoExtsV2(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/certVersion2NoExtensions.pem"
	desEnum := Pass
	out, _ := Lints["e_cert_extensions_version_not_3"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
