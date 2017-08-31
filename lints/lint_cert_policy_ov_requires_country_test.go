// lint_cert_policy_ov_requires_country_test.go
package lints

import (
	"testing"
)

func TestCertPolicyOvHasCountry(t *testing.T) {
	
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_cert_policy_ov_requires_country"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestCertPolicyOvNoCountry(t *testing.T) {
	
	inputPath := "../testlint/testCerts/orgValNoCountry.pem"
	expected := Error
	out := Lints["e_cert_policy_ov_requires_country"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
