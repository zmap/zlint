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
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestPolicyConstraintsCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/policyConstGoodBoth.pem"
	expected := Pass
	out := Lints["e_ext_policy_constraints_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


