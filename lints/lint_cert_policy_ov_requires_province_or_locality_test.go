// lint_cert_policy_ov_requires_province_or_locality_test.go
package lints

import (
	"testing"
)

func TestCertPolicyOvHasCountryOrLocal(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_cert_policy_ov_requires_province_or_locality"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCertPolicyOvNoCountryOrLocal(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/orgValNoProvinceOrLocal.pem"
	expected := Error
	out := Lints["e_cert_policy_ov_requires_province_or_locality"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
