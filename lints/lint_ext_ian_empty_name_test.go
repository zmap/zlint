// lint_ext_ian_empty_name_test.go
package lints

import (
	"testing"
)

func TestIANEmptyName(t *testing.T) {
	inputPath := "../testlint/testCerts/IANEmptyName.pem"
	expected := Error
	out := Lints["e_ext_ian_empty_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestIANNotEmptyName(t *testing.T) {
	inputPath := "../testlint/testCerts/IANDNSIA5String.pem"
	expected := Pass
	out := Lints["e_ext_ian_empty_name"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
