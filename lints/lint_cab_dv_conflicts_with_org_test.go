// lint_cert_policy_conflicts_with_org_test.go
package lints

import (
	"testing"
)

func TestCertPolicyNotConflictWithOrg(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	desEnum := Pass
	out := Lints["e_cab_dv_conflicts_with_org"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCertPolicyConflictsWithOrg(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/domainValWithOrg.pem"
	desEnum := Error
	out := Lints["e_cab_dv_conflicts_with_org"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
