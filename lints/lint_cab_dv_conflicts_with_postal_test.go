package lints

import (
	"testing"
)

func TestCertPolicyNotConflictWithPostal(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	expected := Pass
	out := Lints["e_cab_dv_conflicts_with_postal"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCertPolicyConflictsWithPostal(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValWithPostal.pem"
	expected := Error
	out := Lints["e_cab_dv_conflicts_with_postal"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
