// lint_IAN_IANa_pub_suffix_empty_test.go
package lints

import (
	"testing"
)

func TestIANBarePubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/IANBareSuffix.cer"
	desEnum := Warn
	out, _ := Lints["w_ian_iana_pub_suffix_empty"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIANGoodPubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/IANGoodSuffix.cer"
	desEnum := Pass
	out, _ := Lints["w_ian_iana_pub_suffix_empty"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
