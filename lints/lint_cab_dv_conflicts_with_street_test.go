// lint_cert_policy_conflicts_with_street_test.go
package lints

import (
	"testing"
)

func TestCertPolicyNotConflictWithStreet(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	desEnum := Pass
	out := Lints["e_cab_dv_conflicts_with_street"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCertPolicyConflictsWithStreet(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/domainValWithStreet.pem"
	desEnum := Error
	out := Lints["e_cab_dv_conflicts_with_street"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
