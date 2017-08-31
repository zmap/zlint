package lints

import (
	"testing"
)

func TestSignatureAlgorithmNotSupported(t *testing.T) {
	inputPath := "../testlint/testCerts/md5WithRSASignatureAlgorithm.pem"
	expected := Error
	out := Lints["e_signature_algorithm_not_supported"].Execute(ReadCertificate(inputPath))
	if out.Result != expected {
		t.Error(
			"For", inputPath,
			"expected", expected,
			"got", out.Result,
		)
	}
}

func TestSignatureAlgorithmSHA1Supported(t *testing.T) {
	inputPath := "../testlint/testCerts/sha1WithRSASignatureAlgorithm.pem"
	expected := Pass
	out := Lints["e_signature_algorithm_not_supported"].Execute(ReadCertificate(inputPath))
	if out.Result != expected {
		t.Error(
			"For", inputPath,
			"expected", expected,
			"got", out.Result,
		)
	}
}
