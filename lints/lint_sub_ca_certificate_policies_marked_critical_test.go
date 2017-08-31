// lint_sub_ca_certificate_policies_marked_critical_test.go
package lints

import (
	"testing"
)

func TestSubCaPolicyCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWCertPolicyCrit.pem"
	expected := Warn
	out := Lints["w_sub_ca_certificate_policies_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCaPolicyNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWCertPolicyNoCrit.pem"
	expected := Pass
	out := Lints["w_sub_ca_certificate_policies_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
