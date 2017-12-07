// lint_subject_surname_max_length_test.go
package lints

import (
	"testing"
)

func TestSubjectStreetAddressLengthOK(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectStreetAddress.pem"
	expected := Pass
	out := Lints["e_subject_street_address_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubjectStreetAddressTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectStreetAddressTooLong.pem"
	expected := Error
	out := Lints["e_subject_street_address_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
