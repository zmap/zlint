// lint_ext_san_other_name_present_test.go
package lints

import (
	"testing"
)

func TestSANEmailPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/SANRFC822Beginning.pem"
	expected := Error
	out := Lints["e_ext_san_rfc822_name_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANEmailPresent2(t *testing.T) {
	inputPath := "../testlint/testCerts/SANRFC822End.pem"
	expected := Error
	out := Lints["e_ext_san_rfc822_name_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANEmailMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCaGood.pem"
	expected := Pass
	out := Lints["e_ext_san_rfc822_name_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
