// lint_ext_duplicate_extension_test.go
package lints

import (
	"testing"
)

func TestDuplicateExtension(t *testing.T) {
	inputPath := "../testlint/testCerts/extSANDuplicated.pem"
	expected := Error
	out := Lints["e_ext_duplicate_extension"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestNoDuplicateExtension(t *testing.T) {
	inputPath := "../testlint/testCerts/caBasicConstCrit.pem"
	expected := Pass
	out := Lints["e_ext_duplicate_extension"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
