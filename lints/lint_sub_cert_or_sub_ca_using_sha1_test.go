// lint_sub_cert_or_sub_ca_using_sha1_test.go
package lints

import (
	"testing"
)

// As a note, these certificates were not built, but instead grabbed from censys.io/query
// using the following query to find the raw data and match it to validity period

func TestSHA1After2016(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/RSAWithSHA1after2016.pem"
	desEnum := Error
	out, _ := Lints["e_sub_cert_or_sub_ca_using_sha1"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSHA1Before2016(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/RSAWithSHA1before2016.pem"
	desEnum := Pass
	out, _ := Lints["e_sub_cert_or_sub_ca_using_sha1"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
