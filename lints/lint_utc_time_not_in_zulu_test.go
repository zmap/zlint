// lint_utc_time_not_in_zulu_test.go
package lints

import (

	"testing"
)

func TestUtcZulu(t *testing.T) {
	inputPath := "../testlint/testCerts/utcHasSeconds.cer"
	desEnum := Pass
	out, _ := Lints["utc_time_not_in_zulu"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestUtcNotZulu(t *testing.T) {
	inputPath := "../testlint/testCerts/utcNotZulu.cer"
	desEnum := Error
	out, _ := Lints["utc_time_not_in_zulu"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
