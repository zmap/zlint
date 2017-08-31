// lint_sub_cert_key_usage_cert_sign_bit_set_test.go
package lints

import (
	"testing"
)

func TestGivenNameCorrectPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/givenNameCorrectPolicy.pem"
	expected := Pass
	out := Lints["e_sub_cert_given_name_surname_contains_correct_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSurnameCorrectPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/surnameCorrectPolicy.pem"
	expected := Pass
	out := Lints["e_sub_cert_given_name_surname_contains_correct_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestGivenNameIncorrectPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/givenNameIncorrectPolicy.pem"
	expected := Error
	out := Lints["e_sub_cert_given_name_surname_contains_correct_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSurnameIncorrectPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/surnameIncorrectPolicy.pem"
	expected := Error
	out := Lints["e_sub_cert_given_name_surname_contains_correct_policy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
