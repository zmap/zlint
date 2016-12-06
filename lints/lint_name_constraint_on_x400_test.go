// lint_name_constraint_on_edi_party_name_test.go
package lints

import (

	"testing"
)

func TestNcNoX400(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinZero.cer"
	desEnum := Pass
	out, _ := Lints["name_constraint_on_x400"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestNcX400(t *testing.T) {
	inputPath := "../testlint/testCerts/ncOnX400.cer"
	desEnum := Warn
	out, _ := Lints["name_constraint_on_x400"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
