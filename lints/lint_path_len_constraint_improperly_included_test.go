// lint_path_len_constraint_improperly_included_test.go
package lints

import (
	"testing"
)

func TestCaMaxLenPresentNoCertSign(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathLenPresentNoCertSign.pem"
	expected := Error
	out := Lints["e_path_len_constraint_improperly_included"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaMaxLenPresentGood(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathLenPositive.pem"
	expected := Pass
	out := Lints["e_path_len_constraint_improperly_included"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaMaxLenMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathLenMissing.pem"
	expected := Pass
	out := Lints["e_path_len_constraint_improperly_included"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertMaxLenPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPathLenPositive.pem"
	expected := Error
	out := Lints["e_path_len_constraint_improperly_included"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertMaxLenNone(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_path_len_constraint_improperly_included"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


