// lint_san_iana_pub_suffix_empty_test.go
package lints

import (

	"testing"
)

func TestSanBarePubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/sanBareSuffix.cer"
	desEnum := Warn
	out, _ := Lints["san_iana_pub_suffix_empty"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanGoodPubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/sanGoodSuffix.cer"
	desEnum := Pass
	out, _ := Lints["san_iana_pub_suffix_empty"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
