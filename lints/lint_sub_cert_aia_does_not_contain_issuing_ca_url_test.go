// lint_sub_cert_aia_does_not_contain_issuing_ca_url_test.go
package lints

import (
	"testing"
)

func TestSubCertNoIssuerUrl(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertWOcspURL.pem"
	desEnum := Warn
	out, _ := Lints["e_sub_cert_aia_does_not_contain_issuing_ca_url"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertHasIssuerUrl(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertWIssuerURL.pem"
	desEnum := Pass
	out, _ := Lints["e_sub_cert_aia_does_not_contain_issuing_ca_url"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
