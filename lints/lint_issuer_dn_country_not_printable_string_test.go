// lint_issuer_dn_country_not_printable_string_test.go
package lints

import "testing"

func TestIssuerCountryGood(t *testing.T) {
	inputPath := "../testlint/testCerts/SubjectDNAndIssuerDNCountryPrintableString.pem"
	expected := Pass

	out := Lints["e_issuer_dn_country_not_printable_string"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIssuerCountryBad(t *testing.T) {
	inputPath := "../testlint/testCerts/IssuerDNCountryNotPrintableString.pem"
	expected := Error

	out := Lints["e_issuer_dn_country_not_printable_string"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
