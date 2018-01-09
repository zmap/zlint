package lints

import (
	"testing"
)

func TestSubjectDNLeadingSpace(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectDNLeadingSpace.pem"
	expected := Warn
	out := Lints["w_subject_dn_leading_whitespace"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubjectDNGood(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	expected := Pass
	out := Lints["w_subject_dn_leading_whitespace"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
