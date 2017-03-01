// lint_sub_ca_crl_distribution_points_does_not_contain_url_test.go
package lints

import (
	"testing"
)

func TestSubCaCrlNoUrl(t *testing.T) {
	inputPath := "../testlint/testCerts/subCaCrlMissing.cer"
	desEnum := Error
	out, _ := Lints["sub_ca_crl_distribution_points_does_not_contain_url"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSubCaCrlUrlPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/subCaCrlPresent.cer"
	desEnum := Pass
	out, _ := Lints["sub_ca_crl_distribution_points_does_not_contain_url"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
