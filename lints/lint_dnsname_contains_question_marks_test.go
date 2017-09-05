package lints

import (
	"testing"
)

func TestDNSNameContainsQuestionMark(t *testing.T) {
	inputPath := "../testlint/testCerts/dnsNameContainsQuestionMark.pem"
	expected := Notice
	out := Lints["n_dnsname_contains_question_marks"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
