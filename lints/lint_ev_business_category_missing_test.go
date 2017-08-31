package lints

import (
	"testing"
)

func TestEvNoBiz(t *testing.T) {
	inputPath := "../testlint/testCerts/evAllGood.pem"
	desEnum := Error
	out := Lints["e_ev_business_category_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
