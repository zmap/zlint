// lint_name_constraint_empty_test.go
package lints

import (
	"testing"
)

func TestNoNameConstraint(t *testing.T) {
	inputPath := "../testlint/testCerts/noNameConstraint.pem"
	desEnum := Error
	out, _ := Lints["e_name_constraint_empty"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestHasNameConstraint(t *testing.T) {
	inputPath := "../testlint/testCerts/yesNameConstraint.pem"
	desEnum := Pass
	out, _ := Lints["e_name_constraint_empty"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
