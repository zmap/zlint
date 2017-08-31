// lint_name_constraint_empty_test.go
package lints

import (
	"testing"
)

func TestNoNameConstraint(t *testing.T) {
	inputPath := "../testlint/testCerts/noNameConstraint.pem"
	expected := Error
	out := Lints["e_name_constraint_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestHasNameConstraint(t *testing.T) {
	inputPath := "../testlint/testCerts/yesNameConstraint.pem"
	expected := Pass
	out := Lints["e_name_constraint_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
