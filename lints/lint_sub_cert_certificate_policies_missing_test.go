// lint_sub_cert_certificate_policies_missing_test.go
package lints

import (
	"testing"
)

func TestSubCertPolicyMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPolicyMissing.pem"
	desEnum := Error
	out := Lints["e_sub_cert_certificate_policies_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCertPolicyPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPolicyNoCrit.pem"
	desEnum := Pass
	out := Lints["e_sub_cert_certificate_policies_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
