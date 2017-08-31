// lint_sub_cert_aia_does_not_contain_ocsp_url_test.go
package lints

import (
	"testing"
)

func TestSubCertNoIssuerOcsp(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertWIssuerURL.pem"
	expected := Error
	out := Lints["e_sub_cert_aia_does_not_contain_ocsp_url"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertHasIssuerOcsp(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertWOcspURL.pem"
	expected := Pass
	out := Lints["e_sub_cert_aia_does_not_contain_ocsp_url"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
