// lint_br_san_wildcard_not_first_test.go
package lints

import (
	"testing"
)

func TestBrSANWildcardFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWildcardFirst.pem"
	expected := Error
	out := Lints["e_san_wildcard_not_first"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestBrSANWildcardNotFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIValid.pem"
	expected := Pass
	out := Lints["e_san_wildcard_not_first"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


