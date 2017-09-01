// lint_cert_policy_conflicts_with_org_test.go
package lints

import (
	"testing"
)

func TestCertPolicyNotConflictWithOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	expected := Pass
	out := Lints["e_cab_dv_conflicts_with_org"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCertPolicyConflictsWithOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValWithOrg.pem"
	expected := Error
	out := Lints["e_cab_dv_conflicts_with_org"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
