// lint_sub_ca_crl_distribution_points_does_not_contain_url_test.go
package lints

import (
	"testing"
)

func TestSubCaCrlNoUrl(t *testing.T) {
	inputPath := "../testlint/testCerts/subCaCrlMissing.pem"
	desEnum := Error
	out := Lints["e_sub_ca_crl_distribution_points_does_not_contain_url"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCaCrlUrlPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCaCrlPresent.pem"
	desEnum := Pass
	out := Lints["e_sub_ca_crl_distribution_points_does_not_contain_url"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
