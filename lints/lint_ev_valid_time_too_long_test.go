package lints

import (
	"testing"
)

func TestEvValidTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/evValidTooLong.pem"
	desEnum := Error
	out, _ := Lints["e_ev_valid_time_too_long"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestEvValidNotTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/evValidNotTooLong.pem"
	desEnum := Pass
	out, _ := Lints["e_ev_valid_time_too_long"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
