// lint_cert_policy_conflicts_with_province_test.go
package lints

import (
	"testing"
)

func TestCertPolicyNotConflictWithProv(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	expected := Pass
	out := Lints["e_cab_dv_conflicts_with_province"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCertPolicyConflictsWithProv(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValWithProvince.pem"
	expected := Error
	out := Lints["e_cab_dv_conflicts_with_province"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
