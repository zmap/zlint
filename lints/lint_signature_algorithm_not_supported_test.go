package lints

import (
	"testing"
)

func TestSignatureAlgorithmNotSupported(t *testing.T) {
	inputPath := "../testlint/testCerts/md5WithRSASignatureAlgorithm.pem"
	expected := Error
	out := Lints["e_signature_algorithm_not_supported"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSignatureAlgorithmSHA1Supported(t *testing.T) {
	inputPath := "../testlint/testCerts/sha1WithRSASignatureAlgorithm.pem"
	expected := Pass
	out := Lints["e_signature_algorithm_not_supported"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


