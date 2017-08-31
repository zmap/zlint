// lint_sub_cert_certificate_policies_marked_critical_test.go
package lints

import (
	"testing"
)

func TestSubCertPolicyCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPolicyCrit.pem"
	expected := Warn
	out := Lints["w_sub_cert_certificate_policies_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCertPolicyNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPolicyNoCrit.pem"
	expected := Pass
	out := Lints["w_sub_cert_certificate_policies_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
