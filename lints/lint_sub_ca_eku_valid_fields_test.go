// lint_sub_ca_name_constraints_not_critical_test.go
package lints

import (
	"testing"
)

func TestSubCAEKUValidFields(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAEKUValidFields.pem"
	expected := Pass
	out := Lints["n_sub_ca_eku_not_technically_constrained"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCAEKUNotValidFields(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAEKUNotValidFields.pem"
	expected := Notice
	out := Lints["n_sub_ca_eku_not_technically_constrained"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
