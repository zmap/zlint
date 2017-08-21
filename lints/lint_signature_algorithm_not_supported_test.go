package lints

import (
	"testing"
)

func TestSignatureAlgorithmNotSupported(t *testing.T) {
	inputPath := "../testlint/testCerts/md5WithRSASignatureAlgorithm.pem"
	expected := Error
	out, _ := Lints["e_signature_algorithm_not_supported"].ExecuteTest(ReadCertificate(inputPath))
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
	out, _ := Lints["e_signature_algorithm_not_supported"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != expected {
		t.Error(
			"For", inputPath,
			"expected", expected,
			"got", out.Result,
		)
	}
}
