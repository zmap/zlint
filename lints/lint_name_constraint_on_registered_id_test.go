// lint_name_constraint_on_edi_party_name_test.go
package lints

import (
	"testing"
)

func TestNcNoRegId(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinZero.pem"
	desEnum := Pass
	out, _ := Lints["w_name_constraint_on_registered_id"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestNcRegId(t *testing.T) {
	inputPath := "../testlint/testCerts/ncOnRegId.pem"
	desEnum := Warn
	out, _ := Lints["w_name_constraint_on_registered_id"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
