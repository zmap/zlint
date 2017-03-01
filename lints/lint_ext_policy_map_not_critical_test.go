// lint_ext_policy_map_any_policy_test.go
package lints

import (
	"testing"
)

func TestPolicyMapNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapNotCritical.cer"
	desEnum := Warn
	out, _ := Lints["ext_policy_map_not_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestPolicyMapCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapGood.cer"
	desEnum := Pass
	out, _ := Lints["ext_policy_map_not_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
