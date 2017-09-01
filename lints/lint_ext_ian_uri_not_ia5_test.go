//lint_ext_ian_uri_not_ia5_test.go
package lints

import (
	"testing"
)

func TestIANURIIA5(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIIA5String.pem"
	expected := Pass
	out := Lints["e_ext_ian_uri_not_ia5"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestIANURINotIA5(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURINotIA5String.pem"
	expected := Error
	out := Lints["e_ext_ian_uri_not_ia5"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
