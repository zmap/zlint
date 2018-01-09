package lints

import (
	"testing"
)

func TestCertPolicyDuplicated(t *testing.T) {
	inputPath := "../testlint/testCerts/certPolicyDuplicateShort.pem"
	expected := Error
	out := Lints["e_ext_cert_policy_duplicate"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
func TestCertPolicyDuplicatedAssertion(t *testing.T) {
	inputPath := "../testlint/testCerts/certPolicyAssertionDuplicated.pem"
	expected := Error
	out := Lints["e_ext_cert_policy_duplicate"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCertPolicyNotDuplicated(t *testing.T) {
	inputPath := "../testlint/testCerts/certPolicyNoDuplicate.pem"
	expected := Pass
	out := Lints["e_ext_cert_policy_duplicate"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
