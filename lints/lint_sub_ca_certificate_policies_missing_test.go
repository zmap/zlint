// lint_sub_ca_certificate_policies_missing_test.go
package lints

import (
	"testing"
)

func TestSubCaPolicyMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWNoCertPolicy.pem"
	expected := Error
	out := Lints["e_sub_ca_certificate_policies_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCaPolicyPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWCertPolicyNoCrit.pem"
	expected := Pass
	out := Lints["e_sub_ca_certificate_policies_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


