// lint_sub_cert_cert_policy_empty_test.go
package lints

import (
	"testing"
)

func TestCertPolicyMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPolicyMissing.pem"
	expected := Error
	out := Lints["e_sub_cert_cert_policy_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCertPolicyPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPolicyNoCrit.pem"
	expected := Pass
	out := Lints["e_sub_cert_cert_policy_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
