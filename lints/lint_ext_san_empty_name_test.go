// lint_ext_san_empty_name_test.go
package lints

import (

	"testing"
)

func TestSanEmptyName(t *testing.T) {
	inputPath := "../testlint/testCerts/sanEmptyName.cer"
	desEnum := Error
	out, _ := Lints["ext_san_empty_name"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanNotEmptyName(t *testing.T) {
	inputPath := "../testlint/testCerts/sanCaGood.cer"
	desEnum := Pass
	out, _ := Lints["ext_san_empty_name"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
