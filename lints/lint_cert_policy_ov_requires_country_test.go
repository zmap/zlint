// lint_cert_policy_ov_requires_country_test.go
package lints

import (
	"testing"
)

func TestCertPolicyOvHasCountry(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	desEnum := Pass
	out := Lints["e_cert_policy_ov_requires_country"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCertPolicyOvNoCountry(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/orgValNoCountry.pem"
	desEnum := Error
	out := Lints["e_cert_policy_ov_requires_country"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
