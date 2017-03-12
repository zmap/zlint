// lint_ext_san_other_name_present_test.go
package lints

import (
	"testing"
)

func TestSanEdiPartyPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/sanEdiParty.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_edi_party_name_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanEdiPartyMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/sanOtherName.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_edi_party_name_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
