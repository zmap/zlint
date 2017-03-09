// lint_name_constraint_on_edi_party_name_test.go
package lints

import (
	"testing"
)

func TestNcNoEdi(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinZero.cer"
	desEnum := Pass
	out, _ := Lints["w_name_constraint_on_edi_party_name"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestNcEdi(t *testing.T) {
	inputPath := "../testlint/testCerts/ncOnEdi.cer"
	desEnum := Warn
	out, _ := Lints["w_name_constraint_on_edi_party_name"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
