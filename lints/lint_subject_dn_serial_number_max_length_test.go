// lint_subject_dn_serial_number_max_length_test.go
package lints

import "testing"

func TestSubjectDNSerialNumberBelowMaximumLengthGood(t *testing.T) {
	inputPath := "../testlint/testCerts/evAllGood.pem"
	expected := Pass

	out := Lints["e_subject_dn_serial_number_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubjectDNSerialNumberTooLongBad(t *testing.T) {
	inputPath := "../testlint/testCerts/SubjectDNSerialNumberTooLong.pem"
	expected := Error

	out := Lints["e_subject_dn_serial_number_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
