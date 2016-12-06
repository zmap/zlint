// lint_distribution_point_incomplete_test.go
package lints

import (

	"testing"
)

func crlCompleteDp(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/crlComlepteDp.cer"
	desEnum := Pass
	out, _ := Lints["distribution_point_incomplete"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func crlIncompleteDp(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/crlIncomlepteDp.cer"
	desEnum := Error
	out, _ := Lints["distribution_point_incomplete"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
