// lint_cert_policy_iv_requires_country_test.go
package lints

import (
	"testing"
)

func TestCertPolicyIvCountry(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/indivValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_cert_policy_iv_requires_country"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCertPolicyIvNoCountry(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/indivValNoCountry.pem"
	expected := Error
	out := Lints["e_cert_policy_iv_requires_country"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
