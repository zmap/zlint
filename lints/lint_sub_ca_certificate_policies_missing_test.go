// lint_sub_ca_certificate_policies_missing_test.go
package lints

import (
	"testing"
)

func TestSubCaPolicyMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWNoCertPolicy.pem"
	expected := Error
	out := Lints["e_sub_ca_certificate_policies_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCaPolicyPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWCertPolicyNoCrit.pem"
	expected := Pass
	out := Lints["e_sub_ca_certificate_policies_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
