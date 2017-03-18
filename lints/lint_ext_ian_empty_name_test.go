// lint_ext_ian_empty_name_test.go
package lints

import (
	"testing"
)

func TestIanEmptyName(t *testing.T) {
	inputPath := "../testlint/testCerts/IANEmptyName.cer"
	desEnum := Error
	out, _ := Lints["e_ext_ian_empty_name"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIanNotEmptyName(t *testing.T) {
	inputPath := "../testlint/testCerts/IANDnsIa5.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_ian_empty_name"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
