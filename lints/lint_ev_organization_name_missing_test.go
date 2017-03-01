package lints

import (
	"testing"
)

func TestEvHasOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/evAllGood.cer"
	desEnum := Pass
	out, _ := Lints["ev_organization_name_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestEvNoOrg(t *testing.T) {
	inputPath := "../testlint/testCerts/evNoOrg.cer"
	desEnum := Error
	out, _ := Lints["ev_organization_name_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
