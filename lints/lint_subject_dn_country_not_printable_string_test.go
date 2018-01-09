package lints

import "testing"

func TestSubjectCountryGood(t *testing.T) {
	inputPath := "../testlint/testCerts/SubjectDNAndIssuerDNCountryPrintableString.pem"
	expected := Pass

	out := Lints["e_subject_dn_country_not_printable_string"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubjectCountryBad(t *testing.T) {
	inputPath := "../testlint/testCerts/SubjectDNCountryNotPrintableString.pem"
	expected := Error

	out := Lints["e_subject_dn_country_not_printable_string"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
