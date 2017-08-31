// lint_ext_cert_policy_explicit_text_ia5_string_test.go
package lints

import (
	"testing"
)

func TestExplicitTextIA5String(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticePres.pem"
	expected := Error
	out := Lints["e_ext_cert_policy_explicit_text_ia5_string"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestExplicitTextNotIA5String(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticeExpTextNotIA5String.pem"
	expected := Pass
	out := Lints["e_ext_cert_policy_explicit_text_ia5_string"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestExplicitTextNotPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticeMissing.pem"
	expected := NA
	out := Lints["e_ext_cert_policy_explicit_text_ia5_string"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestExplicitTextNotPresent2(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/userNoticeUnrecommended.pem"
	expected := NA
	out := Lints["e_ext_cert_policy_explicit_text_ia5_string"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
