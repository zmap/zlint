package lints

import (
	"testing"
)

func TestSubEmptyNoSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectEmptyNoSAN.pem"
	expected := Error
	out := Lints["e_subject_empty_without_san"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubEmptyYesSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/SANSubjectEmptyNotCritical.pem"
	expected := Pass
	out := Lints["e_subject_empty_without_san"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
