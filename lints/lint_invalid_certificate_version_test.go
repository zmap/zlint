// lint_invalid_certificate_version_test.go
package lints

import (

	"testing"
)

func TestCertVersion2(t *testing.T) {
	inputPath := "../testlint/testCerts/certVersion2WithExtension.cer"
	desEnum := Error
	out, _ := Lints["invalid_certificate_version"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCertVersion3(t *testing.T) {
	inputPath := "../testlint/testCerts/certVersion3NoExtensions.cer"
	desEnum := Pass
	out, _ := Lints["invalid_certificate_version"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
