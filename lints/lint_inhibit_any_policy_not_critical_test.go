// lint_inhibit_any_policy_not_critical_test.go
package lints

import (
	"testing"
)

func TestInhibitAnyPolicyNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/inhibitAnyNotCrit.pem"
	desEnum := Error
	out, _ := Lints["e_inhibit_any_policy_not_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestInhibitAnyPolicyCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/inhibitAnyCrit.pem"
	desEnum := Pass
	out, _ := Lints["e_inhibit_any_policy_not_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
