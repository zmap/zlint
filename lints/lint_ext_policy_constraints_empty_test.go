// lint_ext_policy_constraints_empty_test.go
package lints

import (
	"testing"
)

func TestPolicyConstraintsEmpty(t *testing.T) {
	inputPath := "../testlint/testCerts/policyConstEmpty.pem"
	desEnum := Error
	out := Lints["e_ext_policy_constraints_empty"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestPolicyConstraintsNotEmpty(t *testing.T) {
	inputPath := "../testlint/testCerts/policyConstGoodBoth.pem"
	desEnum := Pass
	out := Lints["e_ext_policy_constraints_empty"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
