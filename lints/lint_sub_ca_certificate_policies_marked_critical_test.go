// lint_sub_ca_certificate_policies_marked_critical_test.go
package lints

import (
	"testing"
)

func TestSubCaPolicyCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWCertPolicyCrit.pem"
	desEnum := Warn
	out := Lints["w_sub_ca_certificate_policies_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCaPolicyNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWCertPolicyNoCrit.pem"
	desEnum := Pass
	out := Lints["w_sub_ca_certificate_policies_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
