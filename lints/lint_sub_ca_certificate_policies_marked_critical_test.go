package lints

import (
	"testing"
)

func TestSubCaPolicyCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWCertPolicyCrit.pem"
	expected := Warn
	out := Lints["w_sub_ca_certificate_policies_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCaPolicyNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWCertPolicyNoCrit.pem"
	expected := Pass
	out := Lints["w_sub_ca_certificate_policies_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
