// lint_ext_san_missing_test.go
package lints

import (
	"testing"
)

func TestNoSan(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectEmptyNoSan.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestHasSan(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
