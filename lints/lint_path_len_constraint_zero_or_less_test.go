// lint_path_len_constraint_zero_or_less_test.go
package lints

import (
	"testing"
)

func TestCaMaxLenNegative(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathNegative.pem"
	desEnum := Error
	out, _ := Lints["e_path_len_constraint_zero_or_less"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCerMaxLenNegative(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPathLenNegative.pem"
	desEnum := Error
	out, _ := Lints["e_path_len_constraint_zero_or_less"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCaMaxLenPositive(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathLenPositive.pem"
	desEnum := Pass
	out, _ := Lints["e_path_len_constraint_zero_or_less"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertMaxLenPositive(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPathLenPositive.pem"
	desEnum := Pass
	out, _ := Lints["e_path_len_constraint_zero_or_less"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertMaxLenMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/caBasicConstMissing.pem"
	desEnum := NA
	out, _ := Lints["e_path_len_constraint_zero_or_less"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCAMaxLenNone(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathLenMissing.pem"
	desEnum := Pass
	out, _ := Lints["e_path_len_constraint_zero_or_less"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
