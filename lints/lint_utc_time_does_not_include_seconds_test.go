// lint_utc_time_does_not_include_seconds_test.go
package lints

import (
	"testing"
)

func TestUtcHasSeconds(t *testing.T) {
	inputPath := "../testlint/testCerts/utcHasSeconds.pem"
	expected := Pass
	out := Lints["e_utc_time_does_not_include_seconds"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestUtcNoSeconds(t *testing.T) {
	inputPath := "../testlint/testCerts/utcNoSeconds.pem"
	expected := Error
	out := Lints["e_utc_time_does_not_include_seconds"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
