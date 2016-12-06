// lint_sub_cert_cert_policy_empty_test.go
package lints

import (

	"testing"
)

func TestCertPolicyMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPolicyMissing.cer"
	desEnum := Warn
	out, _ := Lints["sub_cert_cert_policy_empty"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCertPolicyPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPolicyNoCrit.cer"
	desEnum := Pass
	out, _ := Lints["sub_cert_cert_policy_empty"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
