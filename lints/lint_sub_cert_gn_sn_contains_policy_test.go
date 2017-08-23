// lint_sub_cert_key_usage_cert_sign_bit_set_test.go
package lints

import (
	"testing"
)

func TestGivenNameCorrectPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/givenNameCorrectPolicy.pem"
	desEnum := Pass
	out, _ := Lints["e_sub_cert_givename_surname_contains_correct_policy_id"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSurnameCorrectPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/surnameCorrectPolicy.pem"
	desEnum := Pass
	out, _ := Lints["e_sub_cert_givename_surname_contains_correct_policy_id"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestGivenNameIncorrectPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/givenNameIncorrectPolicy.pem"
	desEnum := Error
	out, _ := Lints["e_sub_cert_givename_surname_contains_correct_policy_id"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSurnameIncorrectPolicy(t *testing.T) {
	inputPath := "../testlint/testCerts/surnameIncorrectPolicy.pem"
	desEnum := Error
	out, _ := Lints["e_sub_cert_givename_surname_contains_correct_policy_id"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
