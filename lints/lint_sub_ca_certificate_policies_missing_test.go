// lint_sub_ca_certificate_policies_missing_test.go
package lints

import (
	"testing"
)

func TestSubCaPolicyMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWNoCertPolicy.cer"
	desEnum := Error
	out, _ := Lints["e_sub_ca_certificate_policies_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCaPolicyPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWCertPolicyNoCrit.cer"
	desEnum := Pass
	out, _ := Lints["e_sub_ca_certificate_policies_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
