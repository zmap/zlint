// lint_ext_san_other_name_present_test.go
package lints

import (
	"testing"
)

func TestSANEDIPartyPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/SANEDIParty.pem"
	desEnum := Error
	out := Lints["e_ext_san_edi_party_name_present"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSANEDIPartyMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/SANOtherName.pem"
	desEnum := Pass
	out := Lints["e_ext_san_edi_party_name_present"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
