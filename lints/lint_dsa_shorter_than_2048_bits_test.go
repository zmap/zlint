package lints

import "testing"

func TestDSAShorterThan2048Bits(t *testing.T) {
	inputPath := "../testlint/testCerts/dsaShorterThan2048Bits.pem"
	expected := Error
	out := Lints["e_dsa_shorter_than_2048_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestDSANotShorterThan2048Bits(t *testing.T) {
	inputPath := "../testlint/testCerts/dsaNotShorterThan2048Bits.pem"
	expected := Pass
	out := Lints["e_dsa_shorter_than_2048_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
