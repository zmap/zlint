// lint_cert_policy_conflicts_with_street_test.go
package lints

import (
	"testing"
)

func TestCertPolicyNotConflictWithStreet(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/domainValGoodSubject.cer"
	desEnum := Pass
	out, _ := Lints["cert_policy_conflicts_with_street"].ExecuteTest(ReadCertificate(inputPath))
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
	inputPath := "../testlint/testCerts/domainValWithStreet.cer"
	desEnum := Error
	out, _ := Lints["cert_policy_conflicts_with_street"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
