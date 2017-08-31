// lint_cert_policy_requires_org_test.go
package lints

import (
	"testing"
)

func TestCertPolicyOvHasOrg(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	desEnum := Pass
	out := Lints["e_cab_ov_requires_org"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCertPolicyOvNoOrg(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/orgValNoOrg.pem"
	desEnum := Error
	out := Lints["e_cab_ov_requires_org"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
