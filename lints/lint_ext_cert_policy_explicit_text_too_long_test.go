// lint_ext_cert_policy_explicit_text_ia5_string_test.go
package lints

import (
	"testing"
)

func TestExplicitText200Char(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/explicitText200Char.cer"
	desEnum := Error
	out, _ := Lints["ext_cert_policy_explicit_text_too_long"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestExplicitText7Char(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/userNoticeExpTextUtf8.cer"
	desEnum := Pass
	out, _ := Lints["ext_cert_policy_explicit_text_too_long"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
