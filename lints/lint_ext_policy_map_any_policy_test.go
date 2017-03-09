// lint_ext_policy_map_any_policy_test.go
package lints

import (
	"testing"
)

func TestPolicyMapFromAnyPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapFromAnyPolicy.cer"
	desEnum := Error
	out, _ := Lints["e_ext_policy_map_any_policy"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestPolicyMapToAnyPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapToAnyPolicy.cer"
	desEnum := Error
	out, _ := Lints["e_ext_policy_map_any_policy"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestPolicyMapToNoAnyPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapGood.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_policy_map_any_policy"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
