package lints

import (
	"testing"
)

func TestOldSubCertRsaModSizeSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/oldSubTooSmall.pem"
	expected := Error
	out := Lints["e_old_sub_cert_rsa_mod_less_than_1024_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestOldSubCertRsaModSizeNotSmall(t *testing.T) {
	inputPath := "../testlint/testCerts/oldSubSmall.pem"
	expected := Pass
	out := Lints["e_old_sub_cert_rsa_mod_less_than_1024_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
