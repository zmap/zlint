// lint_cert_policy_conflicts_with_province_test.go
package lints

import (
	"testing"
)

func TestCertPolicyNotConflictWithProv(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	desEnum := Pass
	out := Lints["e_cab_dv_conflicts_with_province"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCertPolicyConflictsWithProv(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/domainValWithProvince.pem"
	desEnum := Error
	out := Lints["e_cab_dv_conflicts_with_province"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
