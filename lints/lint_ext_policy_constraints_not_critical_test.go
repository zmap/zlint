// lint_ext_policy_constraints_not_critical_test.go
package lints

import (
	"testing"
)

func TestPolicyConstraintsNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/policyConstNotCritical.pem"
	desEnum := Error
	out := Lints["e_ext_policy_constraints_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestPolicyConstraintsCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/policyConstGoodBoth.pem"
	desEnum := Pass
	out := Lints["e_ext_policy_constraints_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
