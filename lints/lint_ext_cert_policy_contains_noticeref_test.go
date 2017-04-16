// lint_ext_cert_policy_contains_noticeref_test.go
package lints

import (
	"testing"
)

func TestNoticeRefUsed(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/userNoticePres.pem"
	desEnum := Warn
	out, _ := Lints["w_ext_cert_policy_contains_noticeref"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestNoticeRefNotUsed(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/userNoticeMissing.pem"
	desEnum := Pass
	out, _ := Lints["w_ext_cert_policy_contains_noticeref"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
