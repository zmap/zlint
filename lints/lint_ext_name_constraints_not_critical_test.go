// lint_ext_name_constraints_not_critical_test.go
package lints

import (
	"testing"
)

func TestNameConstraintsNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWNameConstNoCrit.pem"
	desEnum := Error
	out, _ := Lints["e_ext_name_constraints_not_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestNameConstraintsCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWNameConstCrit.pem"
	desEnum := Pass
	out, _ := Lints["e_ext_name_constraints_not_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
