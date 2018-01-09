package lints

import (
	"testing"
)

func TestSubCertCountryNameMustAppear(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertCountryNameMustAppear.pem"
	expected := Error
	out := Lints["e_sub_cert_country_name_must_appear"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
