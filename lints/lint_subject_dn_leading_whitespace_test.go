// lint_subject_dn_leading_whitespace_test.go

package lints

import (
	"testing"
)

func TestSubjectDNLeadingSpace(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectDNLeadingSpace.pem"
	expected := Warn
	out := Lints["w_subject_dn_leading_whitespace"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubjectDNGood(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	expected := Pass
	out := Lints["w_subject_dn_leading_whitespace"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
