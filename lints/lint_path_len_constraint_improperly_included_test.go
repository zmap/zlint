// lint_path_len_constraint_improperly_included_test.go
package lints

import (

	"testing"
)

func TestCaMaxLenPresentNoCertSign(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathLenPresentNoCertSign.cer"
	desEnum := Error
	out, _ := Lints["path_len_constraint_improperly_included"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCaMaxLenPresentGood(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathLenPositive.cer"
	desEnum := Pass
	out, _ := Lints["path_len_constraint_improperly_included"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCaMaxLenMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/caMaxPathLenMissing.cer"
	desEnum := Pass
	out, _ := Lints["path_len_constraint_improperly_included"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertMaxLenPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertPathLenPositive.cer"
	desEnum := Error
	out, _ := Lints["path_len_constraint_improperly_included"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertMaxLenNone(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.cer"
	desEnum := Pass
	out, _ := Lints["path_len_constraint_improperly_included"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
