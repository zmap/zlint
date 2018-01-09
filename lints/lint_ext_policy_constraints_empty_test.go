package lints

import (
	"testing"
)

func TestPolicyConstraintsEmpty(t *testing.T) {
	inputPath := "../testlint/testCerts/policyConstEmpty.pem"
	expected := Error
	out := Lints["e_ext_policy_constraints_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestPolicyConstraintsNotEmpty(t *testing.T) {
	inputPath := "../testlint/testCerts/policyConstGoodBoth.pem"
	expected := Pass
	out := Lints["e_ext_policy_constraints_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
