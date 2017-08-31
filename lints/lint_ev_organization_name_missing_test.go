package lints

import (
	"testing"
)

func TestEvHasOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/evAllGood.pem"
	desEnum := Pass
	out := Lints["e_ev_organization_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestEvNoOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/evNoOrg.pem"
	desEnum := Error
	out := Lints["e_ev_organization_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
