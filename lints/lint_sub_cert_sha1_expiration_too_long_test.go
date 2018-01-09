package lints

import (
	"testing"
)

func TestRsaSha1TooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/sha1ExpireAfter2017.pem"
	expected := Warn
	out := Lints["w_sub_cert_sha1_expiration_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestRsaSha1NotTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/sha1ExpirePrior2017.pem"
	expected := Pass
	out := Lints["w_sub_cert_sha1_expiration_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
