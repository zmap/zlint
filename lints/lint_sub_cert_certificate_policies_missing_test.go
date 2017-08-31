// lint_sub_cert_certificate_policies_missing_test.go
package lints

import (
	"testing"
)

func TestSubCertPolicyMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPolicyMissing.pem"
	expected := Error
	out := Lints["e_sub_cert_certificate_policies_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCertPolicyPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPolicyNoCrit.pem"
	expected := Pass
	out := Lints["e_sub_cert_certificate_policies_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
