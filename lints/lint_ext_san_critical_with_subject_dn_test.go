package lints

import (
	"testing"
)

func TestSANCritWithSubjectDn(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCriticalSubjectUncommonOnly.pem"
	expected := Warn
	out := Lints["w_ext_san_critical_with_subject_dn"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANNotCritWithSubjectDn(t *testing.T) {
	inputPath := "../testlint/testCerts/indivValGoodAllFields.pem"
	expected := Pass
	out := Lints["w_ext_san_critical_with_subject_dn"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
