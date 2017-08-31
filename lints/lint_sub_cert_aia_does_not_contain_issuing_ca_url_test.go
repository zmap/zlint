// lint_sub_cert_aia_does_not_contain_issuing_ca_url_test.go
package lints

import (
	"testing"
)

func TestSubCertNoIssuerUrl(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertWOcspURL.pem"
	expected := Warn
	out := Lints["w_sub_cert_aia_does_not_contain_issuing_ca_url"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCertHasIssuerUrl(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertWIssuerURL.pem"
	expected := Pass
	out := Lints["w_sub_cert_aia_does_not_contain_issuing_ca_url"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
