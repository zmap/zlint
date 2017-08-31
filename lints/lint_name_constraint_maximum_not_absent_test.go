// lint_name_constraint_maximum_not_absent_test.go
package lints

import (
	"testing"
)

func TestNcMaxPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/ncAllPres.pem"
	desEnum := Error
	out := Lints["e_name_constraint_maximum_not_absent"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestNcMinPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinPres.pem"
	desEnum := Pass
	out := Lints["e_name_constraint_maximum_not_absent"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestNcEmptyValue(t *testing.T) {
	inputPath := "../testlint/testCerts/ncEmptyValue.pem"
	desEnum := Pass
	out := Lints["e_name_constraint_maximum_not_absent"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
