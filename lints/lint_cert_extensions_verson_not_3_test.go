// lint_cert_extensions_verson_not_3_test.go
package lints

import (

	"testing"
)

func TestExtsV2(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/certVersion2WithExtension.cer"
	desEnum := Error
	out, _ := Lints["cert_extensions_verson_not_3"].ExecuteTest(ReadCertificate(inputPath))
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
	inputPath := "../testlint/testCerts/caBasicConstCrit.cer"
	desEnum := Pass
	out, _ := Lints["cert_extensions_verson_not_3"].ExecuteTest(ReadCertificate(inputPath))
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
	inputPath := "../testlint/testCerts/certVersion2NoExtensions.cer"
	desEnum := Pass
	out, _ := Lints["cert_extensions_verson_not_3"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
