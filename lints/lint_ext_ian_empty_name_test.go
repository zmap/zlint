// lint_ext_ian_empty_name_test.go
package lints

import (
	"testing"
)

func TestIANEmptyName(t *testing.T) {
	inputPath := "../testlint/testCerts/IANEmptyName.pem"
	desEnum := Error
	out := Lints["e_ext_ian_empty_name"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIANNotEmptyName(t *testing.T) {
	inputPath := "../testlint/testCerts/IANDNSIA5String.pem"
	desEnum := Pass
	out := Lints["e_ext_ian_empty_name"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
