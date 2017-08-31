// lint_cert_policy_requires_personal_name_test.go
package lints

import (
	"testing"
)

func TestCertPolicyIvHasPerson(t *testing.T) {
	inputPath := "../testlint/testCerts/indivValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_cab_iv_requires_personal_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCertPolicyIvHasSurname(t *testing.T) {
	inputPath := "../testlint/testCerts/indivValSurnameOnly.pem"
	expected := Error
	out := Lints["e_cab_iv_requires_personal_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCertPolicyIvHasLastName(t *testing.T) {
	inputPath := "../testlint/testCerts/indivValGivenNameOnly.pem"
	expected := Error
	out := Lints["e_cab_iv_requires_personal_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCertPolicyIvNoPerson(t *testing.T) {
	inputPath := "../testlint/testCerts/indivValNoOrgOrPersonalNames.pem"
	expected := Error
	out := Lints["e_cab_iv_requires_personal_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


