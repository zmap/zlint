package lints

import (
	"testing"
)

func TestLeftLabelWildcardCorrect(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/dnsNameWildcardCorrect.pem"
	desEnum := Pass
	out, _ := Lints["e_dnsname_left_label_wildcard_correct"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestLeftLabelWildcardIncorrect(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/dnsNameWildcardIncorrect.pem"
	desEnum := Error
	out, _ := Lints["e_dnsname_left_label_wildcard_correct"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
