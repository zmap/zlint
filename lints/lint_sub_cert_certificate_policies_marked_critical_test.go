package lints

import (
	"testing"
)

func TestSubCertPolicyCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPolicyCrit.pem"
	expected := Warn
	out := Lints["w_sub_cert_certificate_policies_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertPolicyNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPolicyNoCrit.pem"
	expected := Pass
	out := Lints["w_sub_cert_certificate_policies_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
