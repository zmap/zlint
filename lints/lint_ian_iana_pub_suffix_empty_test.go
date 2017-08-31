// lint_ian_iana_pub_suffix_empty_test.go
package lints

import (
	"testing"
)

func TestIANBarePubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/IANBareSuffix.pem"
	expected := Warn
	out := Lints["w_ian_iana_pub_suffix_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIANGoodPubSuffix(t *testing.T) {
	inputPath := "../testlint/testCerts/IANGoodSuffix.pem"
	expected := Pass
	out := Lints["w_ian_iana_pub_suffix_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
