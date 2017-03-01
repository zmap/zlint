// lint_ian_iana_pub_suffix_empty_test.go
package lints

import (
	"testing"
)

func TestIanBarePubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/ianBareSuffix.cer"
	desEnum := Warn
	out, _ := Lints["ian_iana_pub_suffix_empty"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIanGoodPubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/ianGoodSuffix.cer"
	desEnum := Pass
	out, _ := Lints["ian_iana_pub_suffix_empty"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
