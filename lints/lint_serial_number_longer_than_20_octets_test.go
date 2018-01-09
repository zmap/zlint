package lints

import (
	"testing"
)

func TestSnTooLarge(t *testing.T) {
	inputPath := "../testlint/testCerts/serialNumberLarge.pem"
	expected := Error
	out := Lints["e_serial_number_longer_than_20_octets"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSnNotTooLarge(t *testing.T) {
	inputPath := "../testlint/testCerts/serialNumberValid.pem"
	expected := Pass
	out := Lints["e_serial_number_longer_than_20_octets"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
