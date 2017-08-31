// lint_subject_multiple_attr_in_rdn_test.go

package lints

import (
	"testing"
)

func TestSubjectRDNTwoAttribute(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectRDNTwoAttribute.pem"
	expected := Warn
	out := Lints["w_multiple_subject_rdn"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubjectRDNOneAttribute(t *testing.T) {
	inputPath := "../testlint/testCerts/RSASHA1Good.pem"
	expected := Pass
	out := Lints["w_multiple_subject_rdn"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


