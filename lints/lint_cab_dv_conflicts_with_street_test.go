// lint_cert_policy_conflicts_with_street_test.go
package lints

import (
	"testing"
)

func TestCertPolicyNotConflictWithStreet(t *testing.T) {
	
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	expected := Pass
	out := Lints["e_cab_dv_conflicts_with_street"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestCertPolicyConflictsWithStreet(t *testing.T) {
	
	inputPath := "../testlint/testCerts/domainValWithStreet.pem"
	expected := Error
	out := Lints["e_cab_dv_conflicts_with_street"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
