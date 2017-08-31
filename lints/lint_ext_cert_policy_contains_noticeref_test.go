// lint_ext_cert_policy_contains_noticeref_test.go
package lints

import (
	"testing"
)

func TestNoticeRefUsed(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/userNoticePres.pem"
	expected := Warn
	out := Lints["w_ext_cert_policy_contains_noticeref"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestNoticeRefNotUsed(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/userNoticeMissing.pem"
	expected := Pass
	out := Lints["w_ext_cert_policy_contains_noticeref"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
