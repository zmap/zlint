// lint_utc_time_does_not_include_seconds_test.go
package lints

import (
	"testing"
)

func TestUtcHasSeconds(t *testing.T) {
	inputPath := "../testlint/testCerts/utcHasSeconds.cer"
	desEnum := Pass
	out, _ := Lints["e_utc_time_does_not_include_seconds"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestUtcNoSeconds(t *testing.T) {
	inputPath := "../testlint/testCerts/utcNoSeconds.cer"
	desEnum := Error
	out, _ := Lints["e_utc_time_does_not_include_seconds"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
