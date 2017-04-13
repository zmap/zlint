// lint_ext_name_constraints_not_in_ca_test.go
package lints

import (
	"testing"
)

func TestNameConstraintsNotInCa(t *testing.T) {
	inputPath := "../testlint/testCerts/noNameConstraint.pem"
	desEnum := Error
	out, _ := Lints["e_ext_name_constraints_not_in_ca"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestNameConstraintsInCa(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWNameConstCrit.pem"
	desEnum := Pass
	out, _ := Lints["e_ext_name_constraints_not_in_ca"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
