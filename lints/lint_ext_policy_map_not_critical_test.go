// lint_ext_policy_map_any_policy_test.go
package lints

import (
	"testing"
)

func TestPolicyMapNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapNotCritical.pem"
	expected := Warn
	out := Lints["w_ext_policy_map_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestPolicyMapCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/policyMapGood.pem"
	expected := Pass
	out := Lints["w_ext_policy_map_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
