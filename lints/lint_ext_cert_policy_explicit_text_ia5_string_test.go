// lint_ext_cert_policy_explicit_text_ia5_string_test.go
package lints

import (

	"testing"
)

func TestExplicitTextIa5(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticePres.cer"
	desEnum := Error
	out, _ := Lints["ext_cert_policy_explicit_text_ia5_string"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestExplicitTextNotIa5(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticeExpTextNotIa5.cer"
	desEnum := Pass
	out, _ := Lints["ext_cert_policy_explicit_text_ia5_string"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestExplicitTextNotPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticeMissing.cer"
	desEnum := NA
	out, _ := Lints["ext_cert_policy_explicit_text_ia5_string"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestExplicitTextNotPresent2(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/userNoticeUnrecommended.cer"
	desEnum := NA
	out, _ := Lints["ext_cert_policy_explicit_text_ia5_string"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
