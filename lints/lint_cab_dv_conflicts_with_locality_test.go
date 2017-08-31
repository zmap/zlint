// lint_cert_policy_conflicts_with_locality_test.go
package lints

import (
	"testing"
)

func TestCertPolicyNotConflictWithLocal(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	expected := Pass
	out := Lints["e_cab_dv_conflicts_with_locality"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCertPolicyConflictsWithLocal(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/domainValWithLocal.pem"
	expected := Error
	out := Lints["e_cab_dv_conflicts_with_locality"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
