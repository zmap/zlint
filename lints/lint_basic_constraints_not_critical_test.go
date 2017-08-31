// lint_basic_constraints_not_critical_test.go
package lints

import (
	"testing"
)

func TestBasicConstNotCrit(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caBasicConstNotCrit.pem"
	desEnum := Error
	out := Lints["e_basic_constraints_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBasicConstCrit(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caBasicConstCrit.pem"
	desEnum := Pass
	out := Lints["e_basic_constraints_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
