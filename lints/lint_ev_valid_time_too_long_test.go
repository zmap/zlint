package lints

import (
	"testing"
)

func TestEvValidTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/evValidTooLong.pem"
	desEnum := Error
	out := Lints["e_ev_valid_time_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestEvValidNotTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/evValidNotTooLong.pem"
	desEnum := Pass
	out := Lints["e_ev_valid_time_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
