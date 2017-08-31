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
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExplicitTextNotIA5String(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticeExpTextNotIA5String.pem"
	expected := Pass
	out := Lints["e_ext_cert_policy_explicit_text_ia5_string"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExplicitTextNotPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticeMissing.pem"
	expected := NA
	out := Lints["e_ext_cert_policy_explicit_text_ia5_string"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExplicitTextNotPresent2(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticeUnrecommended.pem"
	expected := NA
	out := Lints["e_ext_cert_policy_explicit_text_ia5_string"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
