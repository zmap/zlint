// lint_cert_policy_requires_personal_name_test.go
package lints

import (
	"testing"
)

func TestCertPolicyIvHasPerson(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/indivValGoodAllFields.cer"
	desEnum := Pass
	out, _ := Lints["e_cert_policy_requires_personal_name"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCertPolicyIvHasSurname(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/indivValSurnameOnly.cer"
	desEnum := Error
	out, _ := Lints["e_cert_policy_requires_personal_name"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCertPolicyIvHasLastName(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/indivValGivenNameOnly.cer"
	desEnum := Error
	out, _ := Lints["e_cert_policy_requires_personal_name"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCertPolicyIvNoPerson(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/indivValNoOrgOrPersonalNames.cer"
	desEnum := Error
	out, _ := Lints["e_cert_policy_requires_personal_name"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
