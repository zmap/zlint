// lint_cert_policy_conflicts_with_locality_test.go
package lints

import (
	"testing"
)

func TestCertPolicyNotConflictWithLocal(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/domainValGoodSubject.cer"
	desEnum := Pass
	out, _ := Lints["cert_policy_conflicts_with_locality"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCertPolicyConflictsWithLocal(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/domainValWithLocal.cer"
	desEnum := Error
	out, _ := Lints["cert_policy_conflicts_with_locality"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
