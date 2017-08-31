// lint_cert_policy_iv_requires_country_test.go
package lints

import (
	"testing"
)

func TestCertPolicyIvCountry(t *testing.T) {
	
	inputPath := "../testlint/testCerts/indivValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_cert_policy_iv_requires_country"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestCertPolicyIvNoCountry(t *testing.T) {
	
	inputPath := "../testlint/testCerts/indivValNoCountry.pem"
	expected := Error
	out := Lints["e_cert_policy_iv_requires_country"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
