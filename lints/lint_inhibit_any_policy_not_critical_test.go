// lint_inhibit_any_policy_not_critical_test.go
package lints

import (
	"testing"
)

func TestInhibitAnyPolicyNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/inhibitAnyNotCrit.pem"
	expected := Error
	out := Lints["e_inhibit_any_policy_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestInhibitAnyPolicyCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/inhibitAnyCrit.pem"
	expected := Pass
	out := Lints["e_inhibit_any_policy_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
