// lint_subject_dn_trailing_whitespace_test.go

package lints

import (
	"testing"
)

func TestSubjectDNTrailingSpace(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectDNTrailingSpace.pem"
	expected := Warn
	out := Lints["w_subject_dn_trailing_whitespace"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubjectDNGood2(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	expected := Pass
	out := Lints["w_subject_dn_trailing_whitespace"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
