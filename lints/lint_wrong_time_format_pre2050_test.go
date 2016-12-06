// lint_wrong_time_format_pre2050_test.go
package lints

import (

	"testing"
)

func TestGeneralizedAfter2050(t *testing.T) {
	inputPath := "../testlint/testCerts/generalizedAfter2050.cer"
	desEnum := Pass
	out, _ := Lints["wrong_time_format_pre2050"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestUTCPrior2050(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.cer"
	desEnum := Pass
	out, _ := Lints["wrong_time_format_pre2050"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestGeneralizedPrior2050(t *testing.T) {
	inputPath := "../testlint/testCerts/generalizedPrior2050.cer"
	desEnum := Error
	out, _ := Lints["wrong_time_format_pre2050"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
