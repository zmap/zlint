// lint_ca_country_name_missing_test.go
package lints

import (
	"testing"
)

func TestCaCountryNameMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/caBlankCountry.pem"
	expected := Error
	out := Lints["e_ca_country_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaCountryNamePresent(t *testing.T) {
	inputPath := "../testlint/testCerts/caValCountry.pem"
	expected := Pass
	out := Lints["e_ca_country_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
