// lint_ian_iana_pub_suffix_empty_test.go
package lints

import (
	"testing"
)

func TestIANBarePubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/IANBareSuffix.pem"
	desEnum := Warn
	out := Lints["w_ian_iana_pub_suffix_empty"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIANGoodPubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/IANGoodSuffix.pem"
	desEnum := Pass
	out := Lints["w_ian_iana_pub_suffix_empty"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
