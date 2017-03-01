// lint_sub_cert_crl_distribution_points_does_not_contain_url_test.go
package lints

import (
	"testing"
)

func TestCrlNoUrl(t *testing.T) {
	inputPath := "../testlint/testCerts/subCrlDistNoURL.cer"
	desEnum := Error
	out, _ := Lints["sub_cert_crl_distribution_points_does_not_contain_url"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCrlContainsUrl(t *testing.T) {
	inputPath := "../testlint/testCerts/subCrlDistURL.cer"
	desEnum := Pass
	out, _ := Lints["sub_cert_crl_distribution_points_does_not_contain_url"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
