package lints

import (
	"testing"
)

func TestEvHasLocality(t *testing.T) {
	inputPath := "../testlint/testCerts/evAllGood.pem"
	desEnum := Pass
	out := Lints["e_ev_locality_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestEvNoLocality(t *testing.T) {
	inputPath := "../testlint/testCerts/evNoLocal.pem"
	desEnum := Error
	out := Lints["e_ev_locality_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
