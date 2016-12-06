// lint_name_constraint_maximum_not_absent_test.go
package lints

import (

	"testing"
)

func TestNcMaxPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/ncAllPres.cer"
	desEnum := Error
	out, _ := Lints["name_constraint_maximum_not_absent"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestNcMinPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinPres.cer"
	desEnum := Pass
	out, _ := Lints["name_constraint_maximum_not_absent"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestNcEmptyValue(t *testing.T) {
	inputPath := "../testlint/testCerts/ncEmptyValue.cer"
	desEnum := Pass
	out, _ := Lints["name_constraint_maximum_not_absent"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
