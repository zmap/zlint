package lints

import (
	"testing"
)

func TestLeftLabelWildcardCorrect(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/dnsNameWildcardCorrect.pem"
	expected := Pass
	out := Lints["e_dnsname_left_label_wildcard_correct"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestLeftLabelWildcardIncorrect(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/dnsNameWildcardIncorrect.pem"
	expected := Error
	out := Lints["e_dnsname_left_label_wildcard_correct"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
