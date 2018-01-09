package lints

import (
	"testing"
)

func TestExplicitTextNotUtf8(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticePres.pem"
	expected := Warn
	out := Lints["w_ext_cert_policy_explicit_text_not_utf8"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExplicitTextNotPresentUtf8(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticeMissing.pem"
	expected := NA
	out := Lints["w_ext_cert_policy_explicit_text_not_utf8"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExplicitTextUtf8(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticeExpTextUtf8.pem"
	expected := Pass
	out := Lints["w_ext_cert_policy_explicit_text_not_utf8"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
