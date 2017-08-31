//lint_ext_ian_uri_format_invalid_invalid
package lints

import (
	"testing"
)

func TestIANURIValid(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.pem"
	expected := Pass
	out := Lints["e_ext_ian_uri_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestIANURINoScheme(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURINoScheme.pem"
	expected := Error
	out := Lints["e_ext_san_uri_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestIANURINoSchemeSpecificPart(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURINoSchemeSpecificPart.pem"
	expected := Error
	out := Lints["e_ext_san_uri_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
