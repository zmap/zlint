// lint_sub_ca_name_constraints_not_critical_test.go
package lints

import (
	"testing"
)

func TestSubCaAnyPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWithAnyPolicy.pem"
	expected := Error
	out := Lints["e_sub_ca_must_not_contain_any_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
