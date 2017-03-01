package lints

import (
	"testing"
)

func TestEvHasSN(t *testing.T) {
	inputPath := "../testlint/testCerts/evAllGood.cer"
	desEnum := Pass
	out, _ := Lints["ev_serial_number_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestEvNoSN(t *testing.T) {
	inputPath := "../testlint/testCerts/evNoSN.cer"
	desEnum := Error
	out, _ := Lints["ev_serial_number_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
