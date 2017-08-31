// lint_br_ian_bare_wildcard_test.go
package lints

import (
	"testing"
)

func TestBrIANBareWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/IANBareWildcard.pem"
	expected := Error
	out := Lints["e_ian_bare_wildcard"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestBrIANNotBareWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.pem"
	expected := Pass
	out := Lints["e_ian_bare_wildcard"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
