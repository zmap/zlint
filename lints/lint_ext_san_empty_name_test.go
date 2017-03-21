// lint_ext_san_empty_name_test.go
package lints

import (
	"testing"
)

func TestSANEmptyName(t *testing.T) {
	inputPath := "../testlint/testCerts/SANEmptyName.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_empty_name"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANNotEmptyName(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCaGood.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_empty_name"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
