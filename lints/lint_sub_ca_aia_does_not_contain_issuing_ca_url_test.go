// lint_sub_ca_aia_does_not_contain_issuing_ca_url_test.go
package lints

import (

	"testing"
)

func TestSubCaAiaNoIssuerUrl(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWOcspURL.cer"
	desEnum := Warn
	out, _ := Lints["sub_ca_aia_does_not_contain_issuing_ca_url"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCaAiaHasIssuerUrl(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWBothURL.cer"
	desEnum := Pass
	out, _ := Lints["sub_ca_aia_does_not_contain_issuing_ca_url"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
