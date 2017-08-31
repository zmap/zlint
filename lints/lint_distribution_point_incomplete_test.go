// lint_distribution_point_incomplete_test.go
package lints

import (
	"testing"
)

func crlCompleteDp(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/crlComlepteDp.pem"
	expected := Pass
	out := Lints["e_distribution_point_incomplete"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func crlIncompleteDp(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/crlIncomlepteDp.pem"
	expected := Error
	out := Lints["e_distribution_point_incomplete"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
