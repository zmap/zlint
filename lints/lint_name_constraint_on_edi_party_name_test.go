// lint_name_constraint_on_edi_party_name_test.go
package lints

import (
	"testing"
)

func TestNcNoEDI(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinZero.pem"
	desEnum := Pass
	out := Lints["w_name_constraint_on_edi_party_name"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestNcEDI(t *testing.T) {
	inputPath := "../testlint/testCerts/ncOnEDI.pem"
	desEnum := Warn
	out := Lints["w_name_constraint_on_edi_party_name"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
