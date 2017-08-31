// lint_br_ian_wildcard_not_first_test.go
package lints

import (
	"testing"
)

func TestBrIANWildcardFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/IANWildcardFirst.pem"
	expected := Error
	out := Lints["e_ian_wildcard_not_first"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestBrIANWildcardNotFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.pem"
	expected := Pass
	out := Lints["e_ian_wildcard_not_first"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


