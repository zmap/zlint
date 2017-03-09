package lints

import (
	"testing"
)

func TestEvHasLocality(t *testing.T) {
	inputPath := "../testlint/testCerts/evAllGood.cer"
	desEnum := Pass
	out, _ := Lints["e_ev_locality_name_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestEvNoLocality(t *testing.T) {
	inputPath := "../testlint/testCerts/evNoLocal.cer"
	desEnum := Error
	out, _ := Lints["e_ev_locality_name_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
