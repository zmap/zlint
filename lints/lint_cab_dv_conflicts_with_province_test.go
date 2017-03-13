// lint_cert_policy_conflicts_with_province_test.go
package lints

import (
	"testing"
)

func TestCertPolicyNotConflictWithProv(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/domainValGoodSubject.cer"
	desEnum := Pass
	out, _ := Lints["e_cab_dv_conflicts_with_province"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCertPolicyConflictsWithProv(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/domainValWithProvince.cer"
	desEnum := Error
	out, _ := Lints["e_cab_dv_conflicts_with_province"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
