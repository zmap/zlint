// lint_san_iana_pub_suffix_empty_test.go
package lints

import (
	"testing"
)

func TestSANBarePubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/SANBareSuffix.cer"
	desEnum := Warn
	out, _ := Lints["w_san_iana_pub_suffix_empty"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANGoodPubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/SANGoodSuffix.cer"
	desEnum := Pass
	out, _ := Lints["w_san_iana_pub_suffix_empty"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
