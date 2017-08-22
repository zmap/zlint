// lint_sub_ca_name_constraints_not_critical_test.go
package lints

import (
	"testing"
)

func TestSubCAEKUValidFields(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAEKUValidFields.pem"
	desEnum := Pass
	out, _ := Lints["n_sub_ca_eku_valid_fields"].ExecuteTest(ReadCertificate(inputPath))
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
	out, _ := Lints["n_sub_ca_eku_valid_fields"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
