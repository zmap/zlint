// lint_br_san_bare_wildcard_test.go
package lints

import (
	"testing"
)

func TestBrSANBareWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/SANBareWildcard.pem"
	expected := Error
	out := Lints["e_san_bare_wildcard"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestBrSANNotBareWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIValid.pem"
	expected := Pass
	out := Lints["e_san_bare_wildcard"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
