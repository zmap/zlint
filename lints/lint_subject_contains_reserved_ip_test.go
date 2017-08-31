// lint_subject_contains_reserved_ip_test.go
package lints

import (
	"testing"
)

func TestSubjectIPReserved(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectReservedIP.pem"
	expected := Error
	out := Lints["e_subject_contains_reserved_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubjectIPReserved6(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectReservedIP6.pem"
	expected := Error
	out := Lints["e_subject_contains_reserved_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubjectIPNotReserved(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectGoodIP.pem"
	expected := Pass
	out := Lints["e_subject_contains_reserved_ip"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


