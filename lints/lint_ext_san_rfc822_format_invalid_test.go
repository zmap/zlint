// lint_ext_san_space_dns_name_test.go
package lints

import (
	"testing"
)

func TestSANInvalidEmail(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWithInvalidEmail.pem"
	expected := Error
	out := Lints["e_ext_san_rfc822_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANInvalidEmail2(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWithInvalidEmail2.pem"
	expected := Error
	out := Lints["e_ext_san_rfc822_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANValidEmail(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWithValidEmail.pem"
	expected := Pass
	out := Lints["e_ext_san_rfc822_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


