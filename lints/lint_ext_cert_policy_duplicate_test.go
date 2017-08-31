// lint_ext_cert_policy_duplicate_test.go
package lints

import (
	"testing"
)

func TestCertPolicyDuplicated(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/certPolicyDuplicateShort.pem"
	expected := Error
	out := Lints["e_ext_cert_policy_duplicate"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
func TestCertPolicyDuplicatedAssertion(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/certPolicyAssertionDuplicated.pem"
	expected := Error
	out := Lints["e_ext_cert_policy_duplicate"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCertPolicyNotDuplicated(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/certPolicyNoDuplicate.pem"
	expected := Pass
	out := Lints["e_ext_cert_policy_duplicate"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
