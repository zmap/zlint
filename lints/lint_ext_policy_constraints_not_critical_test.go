// lint_ext_policy_constraints_not_critical_test.go
package lints

import (
	"testing"
)

func TestPolicyConstraintsNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/policyConstNotCritical.pem"
	expected := Error
	out := Lints["e_ext_policy_constraints_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestPolicyConstraintsCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/policyConstGoodBoth.pem"
	expected := Pass
	out := Lints["e_ext_policy_constraints_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
