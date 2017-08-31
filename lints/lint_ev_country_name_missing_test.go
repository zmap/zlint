package lints

import (
	"testing"
)

func TestEvHasCountry(t *testing.T) {
	inputPath := "../testlint/testCerts/evAllGood.pem"
	desEnum := Pass
	out := Lints["e_ev_country_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestEvNoCountry(t *testing.T) {
	inputPath := "../testlint/testCerts/evNoCountry.pem"
	desEnum := Error
	out := Lints["e_ev_country_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
