//lint_ext_san_uri_format_invalid
package lints

import (
	"testing"
)

func TestSANURIValid(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIValid.pem"
	expected := Pass
	out := Lints["e_ext_san_uri_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANURINoScheme(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURINoScheme.pem"
	expected := Error
	out := Lints["e_ext_san_uri_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANURINoSchemeSpecificPart(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURINoSchemeSpecificPart.pem"
	expected := Error
	out := Lints["e_ext_san_uri_format_invalid"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
