// lint_path_len_constraint_zero_or_less_test.go
package lints

import (
	"testing"
)

func TestCaMaxLenNegative(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathNegative.pem"
	expected := Error
	out := Lints["e_path_len_constraint_zero_or_less"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCerMaxLenNegative(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPathLenNegative.pem"
	expected := Error
	out := Lints["e_path_len_constraint_zero_or_less"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaMaxLenPositive(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathLenPositive.pem"
	expected := Pass
	out := Lints["e_path_len_constraint_zero_or_less"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertMaxLenPositive(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPathLenPositive.pem"
	expected := Pass
	out := Lints["e_path_len_constraint_zero_or_less"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertMaxLenMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/caBasicConstMissing.pem"
	expected := NA
	out := Lints["e_path_len_constraint_zero_or_less"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCAMaxLenNone(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathLenMissing.pem"
	expected := Pass
	out := Lints["e_path_len_constraint_zero_or_less"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
