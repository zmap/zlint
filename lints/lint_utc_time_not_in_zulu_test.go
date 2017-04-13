// lint_utc_time_not_in_zulu_test.go
package lints

import (
	"testing"
)

func TestUtcZulu(t *testing.T) {
	inputPath := "../testlint/testCerts/utcHasSeconds.pem"
	desEnum := Pass
	out, _ := Lints["e_utc_time_not_in_zulu"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestUtcNotZulu(t *testing.T) {
	inputPath := "../testlint/testCerts/utcNotZulu.pem"
	desEnum := Error
	out, _ := Lints["e_utc_time_not_in_zulu"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
