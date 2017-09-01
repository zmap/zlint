// lint_ext_cert_policy_disallowed_any_policy_qualifier_test.go
package lints

import (
	"testing"
)

func TestNoticeRef(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticePres.pem"
	expected := Pass
	out := Lints["e_ext_cert_policy_disallowed_any_policy_qualifier"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCps(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticeMissing.pem"
	expected := Pass
	out := Lints["e_ext_cert_policy_disallowed_any_policy_qualifier"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestNoticeRefUnknown(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticeUnrecommended.pem"
	expected := Error
	out := Lints["e_ext_cert_policy_disallowed_any_policy_qualifier"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
