// lint_cert_policy_iv_requires_province_or_locality_test.go
package lints

import (
	"testing"
)

func TestCertPolicyHasCountryOrLocal(t *testing.T) {
	inputPath := "../testlint/testCerts/indivValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_cert_policy_iv_requires_province_or_locality"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCertPolicyIvNoCountryOrLocal(t *testing.T) {
	inputPath := "../testlint/testCerts/indivValNoLocalOrProvince.pem"
	expected := Error
	out := Lints["e_cert_policy_iv_requires_province_or_locality"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


