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
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCerMaxLenNegative(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPathLenNegative.pem"
	expected := Error
	out := Lints["e_path_len_constraint_zero_or_less"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCaMaxLenPositive(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathLenPositive.pem"
	expected := Pass
	out := Lints["e_path_len_constraint_zero_or_less"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCertMaxLenPositive(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPathLenPositive.pem"
	expected := Pass
	out := Lints["e_path_len_constraint_zero_or_less"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCertMaxLenMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/caBasicConstMissing.pem"
	expected := NA
	out := Lints["e_path_len_constraint_zero_or_less"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCAMaxLenNone(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathLenMissing.pem"
	expected := Pass
	out := Lints["e_path_len_constraint_zero_or_less"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
