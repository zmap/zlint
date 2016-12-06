// lint_sub_cert_certificate_policies_missing_test.go
package lints

import (

	"testing"
)

func TestSubCertPolicyMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPolicyMissing.cer"
	desEnum := Error
	out, _ := Lints["sub_cert_certificate_policies_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertPolicyPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPolicyNoCrit.cer"
	desEnum := Pass
	out, _ := Lints["sub_cert_certificate_policies_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
