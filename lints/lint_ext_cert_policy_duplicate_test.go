// lint_ext_cert_policy_duplicate_test.go
package lints

import (
	"testing"
)

func TestCertPolicyDuplicated(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/certPolicyDuplicateShort.pem"
	desEnum := Error
	out := Lints["e_ext_cert_policy_duplicate"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
func TestCertPolicyDuplicatedAssertion(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/certPolicyAssertionDuplicated.pem"
	desEnum := Error
	out := Lints["e_ext_cert_policy_duplicate"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCertPolicyNotDuplicated(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/certPolicyNoDuplicate.pem"
	desEnum := Pass
	out := Lints["e_ext_cert_policy_duplicate"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
