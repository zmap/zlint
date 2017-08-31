// lint_sub_cert_crl_distribution_points_does_not_contain_url_test.go
package lints

import (
	"testing"
)

func TestCrlNoUrl(t *testing.T) {
	inputPath := "../testlint/testCerts/subCrlDistNoURL.pem"
	expected := Error
	out := Lints["e_sub_cert_crl_distribution_points_does_not_contain_url"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestCrlContainsUrl(t *testing.T) {
	inputPath := "../testlint/testCerts/subCrlDistURL.pem"
	expected := Pass
	out := Lints["e_sub_cert_crl_distribution_points_does_not_contain_url"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
