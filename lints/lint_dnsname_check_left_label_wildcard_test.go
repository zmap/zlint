package lints

import (
	"testing"
)

func TestLeftLabelWildcardCorrect(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/dnsNameWildcardCorrect.pem"
	desEnum := Pass
	out := Lints["e_dnsname_left_label_wildcard_correct"].Execute(ReadCertificate(inputPath))
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
	out := Lints["e_dnsname_left_label_wildcard_correct"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
