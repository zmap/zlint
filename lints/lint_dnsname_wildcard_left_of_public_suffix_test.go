package lints

import (
	"testing"
)

func TestWildcardLeftOfPublicSuffix(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/dnsNameWildcardLeftOfPublicSuffix.pem"
	expected := Warn
	out := Lints["w_dnsname_wildcard_left_of_public_suffix"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestWildcardNotLeftOfPublicSuffix(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/dnsNameWildcardNotLeftOfPublicSuffix.pem"
	expected := Pass
	out := Lints["w_dnsname_wildcard_left_of_public_suffix"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
