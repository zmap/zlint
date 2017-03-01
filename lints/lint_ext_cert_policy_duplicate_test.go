// lint_ext_cert_policy_duplicate_test.go
package lints

import (
	"testing"
)

func TestCertPolicyDuplicated(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/certPolicyDuplicateShort.cer"
	desEnum := Error
	out, _ := Lints["ext_cert_policy_duplicate"].ExecuteTest(ReadCertificate(inputPath))
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
	inputPath := "../testlint/testCerts/certPolicyAssertionDuplicated.cer"
	desEnum := Error
	out, _ := Lints["ext_cert_policy_duplicate"].ExecuteTest(ReadCertificate(inputPath))
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
	inputPath := "../testlint/testCerts/certPolicyNoDuplicate.cer"
	desEnum := Pass
	out, _ := Lints["ext_cert_policy_duplicate"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
