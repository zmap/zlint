// lint_cert_policy_iv_requires_province_or_locality_test.go
package lints

import (
	"testing"
)

func TestCertPolicyHasCountryOrLocal(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/indivValGoodAllFields.cer"
	desEnum := Pass
	out, _ := Lints["cert_policy_iv_requires_province_or_locality"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCertPolicyIvNoCountryOrLocal(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/indivValNoLocalOrProvince.cer"
	desEnum := Error
	out, _ := Lints["cert_policy_iv_requires_province_or_locality"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
