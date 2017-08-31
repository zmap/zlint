// lint_sub_ca_name_constraints_not_critical_test.go
package lints

import (
	"testing"
)

func TestSubCAEKUValidFields(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAEKUValidFields.pem"
	desEnum := Pass
	out := Lints["n_sub_ca_eku_not_technically_constrained"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCAEKUNotValidFields(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAEKUNotValidFields.pem"
	desEnum := Notice
	out := Lints["n_sub_ca_eku_not_technically_constrained"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
