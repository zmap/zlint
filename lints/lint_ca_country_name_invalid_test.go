// lint_ca_country_name_invalid_test.go
package lints

import (
	"testing"
)

func TestCaCountryNameInvalid(t *testing.T) {
	inputPath := "../testlint/testCerts/caInvalCountryCode.pem"
	expected := Error
	out := Lints["e_ca_country_name_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaCountryNameValid(t *testing.T) {
	inputPath := "../testlint/testCerts/caValCountry.pem"
	expected := Pass
	out := Lints["e_ca_country_name_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


