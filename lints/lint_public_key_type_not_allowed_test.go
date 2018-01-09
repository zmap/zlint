package lints

import (
	"testing"
)

func TestPKTypeUnknown(t *testing.T) {
	inputPath := "../testlint/testCerts/unknownpublickey.pem"
	expected := Error
	out := Lints["e_public_key_type_not_allowed"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestPKTypeRSA(t *testing.T) {
	inputPath := "../testlint/testCerts/rsawithsha1before2016.pem"
	expected := Pass
	out := Lints["e_public_key_type_not_allowed"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestPKTypeECDSA(t *testing.T) {
	inputPath := "../testlint/testCerts/ecdsaP256.pem"
	expected := Pass
	out := Lints["e_public_key_type_not_allowed"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
