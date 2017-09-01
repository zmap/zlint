//lint_ext_san_uri_not_ia5_test.go
package lints

import (
	"testing"
)

func TestSANURIIA5(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIIA5.pem"
	expected := Pass
	out := Lints["e_ext_san_uri_not_ia5"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANURINotIA5(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURINotIA5.pem"
	expected := Error
	out := Lints["e_ext_san_uri_not_ia5"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
