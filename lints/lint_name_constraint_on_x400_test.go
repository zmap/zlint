// lint_name_constraint_on_edi_party_name_test.go
package lints

import (
	"testing"
)

func TestNcNoX400(t *testing.T) {
	inputPath := "../testlint/testCerts/ncMinZero.pem"
	desEnum := Pass
	out := Lints["w_name_constraint_on_x400"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestNcX400(t *testing.T) {
	inputPath := "../testlint/testCerts/ncOnX400.pem"
	desEnum := Warn
	out := Lints["w_name_constraint_on_x400"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
