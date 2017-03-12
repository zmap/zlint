package lints

import (
	"testing"
)

func TestEvHasCountry(t *testing.T) {
	inputPath := "../testlint/testCerts/evAllGood.cer"
	desEnum := Pass
	out, _ := Lints["e_ev_country_name_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestEvNoCountry(t *testing.T) {
	inputPath := "../testlint/testCerts/evNoCountry.cer"
	desEnum := Error
	out, _ := Lints["e_ev_country_name_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
