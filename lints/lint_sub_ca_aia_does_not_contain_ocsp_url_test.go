// lint_sub_ca_aia_does_not_contain_ocsp_url_test.go
package lints

import (
	"testing"
)

func TestSubCaAiaNoOcsp(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWIssuerURL.pem"
	expected := Error
	out := Lints["e_sub_ca_aia_does_not_contain_ocsp_url"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCaAiaHasOcsp(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWOcspURL.pem"
	expected := Pass
	out := Lints["e_sub_ca_aia_does_not_contain_ocsp_url"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
