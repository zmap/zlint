// lint_ext_cert_policy_explicit_text_ia5_string_test.go
package lints

import (
	"testing"
)

func TestExplicitText200Char(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/explicitText200Char.pem"
	expected := Error
	out := Lints["e_ext_cert_policy_explicit_text_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestExplicitText7Char(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/userNoticeExpTextUtf8.pem"
	expected := Pass
	out := Lints["e_ext_cert_policy_explicit_text_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
