// lint_ext_san_empty_name_test.go
package lints

import (
	"testing"
)

func TestSANEmptyName(t *testing.T) {
	inputPath := "../testlint/testCerts/SANEmptyName.pem"
	expected := Error
	out := Lints["e_ext_san_empty_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSANNotEmptyName(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCaGood.pem"
	expected := Pass
	out := Lints["e_ext_san_empty_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
