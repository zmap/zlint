// lint_ext_cert_policy_explicit_text_ia5_string_test.go
package lints

import (
	"testing"
)

func TestExplicitText200Char(t *testing.T) {
	inputPath := "../testlint/testCerts/explicitText200Char.pem"
	expected := Error
	out := Lints["e_ext_cert_policy_explicit_text_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExplicitText7Char(t *testing.T) {
	inputPath := "../testlint/testCerts/userNoticeExpTextUtf8.pem"
	expected := Pass
	out := Lints["e_ext_cert_policy_explicit_text_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


