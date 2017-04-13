// lint_subject_country_not_iso_test.go
package lints

import (
	"testing"
)

func TestGenralizedNotZulu(t *testing.T) {
	inputPath := "../testlint/testCerts/generalizedNotZulu.pem"
	desEnum := Error
	out, _ := Lints["e_generalized_time_not_in_zulu"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestGenralizedZulu(t *testing.T) {
	inputPath := "../testlint/testCerts/generalizedHasSeconds.pem"
	desEnum := Pass
	out, _ := Lints["e_generalized_time_not_in_zulu"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
