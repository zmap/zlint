// lint_ext_policy_constraints_empty_test.go
package lints

import (
	"testing"
)

func TestPolicyConstraintsEmpty(t *testing.T) {
	inputPath := "../testlint/testCerts/policyConstEmpty.cer"
	desEnum := Error
	out, _ := Lints["e_ext_policy_constraints_empty"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestPolicyConstraintsNotEmpty(t *testing.T) {
	inputPath := "../testlint/testCerts/policyConstGoodBoth.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_policy_constraints_empty"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
