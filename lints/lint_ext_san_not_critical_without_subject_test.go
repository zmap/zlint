package lints

import (
	"testing"
)

func TestSubjectEmptySANNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/SANSubjectEmptyNotCritical.pem"
	expected := Error
	out := Lints["e_ext_san_not_critical_without_subject"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubjectEmptySANCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCaEmptySubject.pem"
	expected := Pass
	out := Lints["e_ext_san_not_critical_without_subject"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubjectNotEmptySANCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCriticalSubjectUncommonOnly.pem"
	expected := Pass
	out := Lints["e_ext_san_not_critical_without_subject"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
