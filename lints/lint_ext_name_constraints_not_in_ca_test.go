// lint_ext_name_constraints_not_in_ca_test.go
package lints

import (
	"testing"
)

func TestNameConstraintsNotInCa(t *testing.T) {
	inputPath := "../testlint/testCerts/noNameConstraint.pem"
	desEnum := Error
	out := Lints["e_ext_name_constraints_not_in_ca"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestNameConstraintsInCa(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWNameConstCrit.pem"
	desEnum := Pass
	out := Lints["e_ext_name_constraints_not_in_ca"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
