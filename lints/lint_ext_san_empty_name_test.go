// lint_ext_san_empty_name_test.go
package lints

import (
	"testing"
)

func TestSANEmptyName(t *testing.T) {
	inputPath := "../testlint/testCerts/SANEmptyName.pem"
	desEnum := Error
	out := Lints["e_ext_san_empty_name"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANNotEmptyName(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCaGood.pem"
	desEnum := Pass
	out := Lints["e_ext_san_empty_name"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
