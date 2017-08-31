// lint_ext_cert_policy_contains_noticeref_test.go
package lints

import (
	"testing"
)

func TestNoticeRefUsed(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticePres.pem"
	expected := Warn
	out := Lints["w_ext_cert_policy_contains_noticeref"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestNoticeRefNotUsed(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticeMissing.pem"
	expected := Pass
	out := Lints["w_ext_cert_policy_contains_noticeref"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


