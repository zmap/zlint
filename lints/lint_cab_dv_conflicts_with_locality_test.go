package lints

import (
	"testing"
)

func TestCertPolicyNotConflictWithLocal(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	expected := Pass
	out := Lints["e_cab_dv_conflicts_with_locality"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCertPolicyConflictsWithLocal(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValWithLocal.pem"
	expected := Error
	out := Lints["e_cab_dv_conflicts_with_locality"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
