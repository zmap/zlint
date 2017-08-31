// lint_sub_ca_certificate_policies_missing_test.go
package lints

import (
	"testing"
)

func TestSubCaPolicyMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWNoCertPolicy.pem"
	desEnum := Error
	out := Lints["e_sub_ca_certificate_policies_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCaPolicyPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWCertPolicyNoCrit.pem"
	desEnum := Pass
	out := Lints["e_sub_ca_certificate_policies_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
