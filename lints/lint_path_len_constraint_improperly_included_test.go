// lint_path_len_constraint_improperly_included_test.go
package lints

import (
	"testing"
)

func TestCaMaxLenPresentNoCertSign(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathLenPresentNoCertSign.pem"
	desEnum := Error
	out := Lints["e_path_len_constraint_improperly_included"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCaMaxLenPresentGood(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathLenPositive.pem"
	desEnum := Pass
	out := Lints["e_path_len_constraint_improperly_included"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCaMaxLenMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathLenMissing.pem"
	desEnum := Pass
	out := Lints["e_path_len_constraint_improperly_included"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertMaxLenPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPathLenPositive.pem"
	desEnum := Error
	out := Lints["e_path_len_constraint_improperly_included"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertMaxLenNone(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	desEnum := Pass
	out := Lints["e_path_len_constraint_improperly_included"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
