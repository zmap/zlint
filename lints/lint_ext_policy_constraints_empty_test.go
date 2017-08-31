// lint_ext_policy_constraints_empty_test.go
package lints

import (
	"testing"
)

func TestPolicyConstraintsEmpty(t *testing.T) {
	inputPath := "../testlint/testCerts/policyConstEmpty.pem"
	expected := Error
	out := Lints["e_ext_policy_constraints_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestPolicyConstraintsNotEmpty(t *testing.T) {
	inputPath := "../testlint/testCerts/policyConstGoodBoth.pem"
	expected := Pass
	out := Lints["e_ext_policy_constraints_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
