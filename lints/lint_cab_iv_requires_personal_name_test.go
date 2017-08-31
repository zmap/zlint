// lint_cert_policy_requires_personal_name_test.go
package lints

import (
	"testing"
)

func TestCertPolicyIvHasPerson(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/indivValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_cab_iv_requires_personal_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCertPolicyIvHasSurname(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/indivValSurnameOnly.pem"
	expected := Error
	out := Lints["e_cab_iv_requires_personal_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCertPolicyIvHasLastName(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/indivValGivenNameOnly.pem"
	expected := Error
	out := Lints["e_cab_iv_requires_personal_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCertPolicyIvNoPerson(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/indivValNoOrgOrPersonalNames.pem"
	expected := Error
	out := Lints["e_cab_iv_requires_personal_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
