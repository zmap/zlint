// lint_sub_cert_aia_does_not_contain_ocsp_url_test.go
package lints

import (
	"testing"
)

func TestSubCertNoIssuerOcsp(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertWIssuerURL.cer"
	desEnum := Error
	out, _ := Lints["e_sub_cert_aia_does_not_contain_ocsp_url"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCertHasIssuerOcsp(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertWOcspURL.cer"
	desEnum := Pass
	out, _ := Lints["e_sub_cert_aia_does_not_contain_ocsp_url"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
