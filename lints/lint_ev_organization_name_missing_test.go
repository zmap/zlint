package lints

import (
	"testing"
)

func TestEvHasOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/evAllGood.pem"
	expected := Pass
	out := Lints["e_ev_organization_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestEvNoOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/evNoOrg.pem"
	expected := Error
	out := Lints["e_ev_organization_name_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
