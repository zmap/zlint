package lints

import (
	"testing"
)

func TestWildcardLeftOfPublicSuffix(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/dnsNameWildcardLeftOfPublicSuffix.pem"
	desEnum := Warn
	out, _ := Lints["e_dnsname_wildcard_left_of_public_suffix"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestWildcardNotLeftOfPublicSuffix(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/dnsNameWildcardNotLeftOfPublicSuffix.pem"
	desEnum := Pass
	out, _ := Lints["e_dnsname_wildcard_left_of_public_suffix"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
