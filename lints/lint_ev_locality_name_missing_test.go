package lints

import (
	"testing"
)

func TestEvHasLocality(t *testing.T) {
	inputPath := "../testlint/testCerts/evAllGood.pem"
	expected := Pass
	out := Lints["e_ev_locality_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestEvNoLocality(t *testing.T) {
	inputPath := "../testlint/testCerts/evNoLocal.pem"
	expected := Error
	out := Lints["e_ev_locality_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
