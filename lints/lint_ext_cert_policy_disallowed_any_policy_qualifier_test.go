// lint_ext_cert_policy_disallowed_any_policy_qualifier_test.go
package lints

import (
	"testing"
)

func TestNoticeRef(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/userNoticePres.pem"
	expected := Pass
	out := Lints["e_ext_cert_policy_disallowed_any_policy_qualifier"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCps(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/userNoticeMissing.pem"
	expected := Pass
	out := Lints["e_ext_cert_policy_disallowed_any_policy_qualifier"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestNoticeRefUnknown(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/userNoticeUnrecommended.pem"
	expected := Error
	out := Lints["e_ext_cert_policy_disallowed_any_policy_qualifier"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
