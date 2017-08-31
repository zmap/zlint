// lint_ext_san_other_name_present_test.go
package lints

import (
	"testing"
)

func TestSANDirNamePresent2(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDirectoryNameBeginning.pem"
	expected := Error
	out := Lints["e_ext_san_directory_name_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANDirNamePresent(t *testing.T) {
	inputPath := "../testlint/testCerts/SANDirectoryNameEnd.pem"
	expected := Error
	out := Lints["e_ext_san_directory_name_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANDirNameMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCaGood.pem"
	expected := Pass
	out := Lints["e_ext_san_directory_name_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
