package lints

import (
	"testing"
)

func TestEvNoBiz(t *testing.T) {
	inputPath := "../testlint/testCerts/evAllGood.cer"
	desEnum := Error
	out, _ := Lints["ev_business_category_missing"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
