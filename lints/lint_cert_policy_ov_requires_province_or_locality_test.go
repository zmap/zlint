// lint_cert_policy_ov_requires_province_or_locality_test.go
package lints

import (
	"testing"
)

func TestCertPolicyOvHasCountryOrLocal(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_cert_policy_ov_requires_province_or_locality"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCertPolicyOvNoCountryOrLocal(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValNoProvinceOrLocal.pem"
	expected := Error
	out := Lints["e_cert_policy_ov_requires_province_or_locality"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
