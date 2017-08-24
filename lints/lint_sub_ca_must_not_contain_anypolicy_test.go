// lint_sub_ca_name_constraints_not_critical_test.go
package lints

import (
	"testing"
)

func TestSubCaAnyPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWithAnyPolicy.pem"
	desEnum := Error
	out, _ := Lints["e_sub_ca_must_not_contain_any_policy"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
