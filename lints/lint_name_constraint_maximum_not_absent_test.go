// lint_name_constraint_maximum_not_absent_test.go
package lints

import (
	"testing"
)

func TestNcMaxPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/ncAllPres.pem"
	expected := Error
	out := Lints["e_name_constraint_maximum_not_absent"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestNcMinPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinPres.pem"
	expected := Pass
	out := Lints["e_name_constraint_maximum_not_absent"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestNcEmptyValue(t *testing.T) {
	inputPath := "../testlint/testCerts/ncEmptyValue.pem"
	expected := Pass
	out := Lints["e_name_constraint_maximum_not_absent"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
