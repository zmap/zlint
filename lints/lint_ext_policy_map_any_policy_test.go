// lint_ext_policy_map_any_policy_test.go
package lints

import (
	"testing"
)

func TestPolicyMapFromAnyPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapFromAnyPolicy.pem"
	expected := Error
	out := Lints["e_ext_policy_map_any_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestPolicyMapToAnyPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapToAnyPolicy.pem"
	expected := Error
	out := Lints["e_ext_policy_map_any_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestPolicyMapToNoAnyPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapGood.pem"
	expected := Pass
	out := Lints["e_ext_policy_map_any_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
