// lint_sub_cert_key_usage_crl_sign_bit_set_test.go
package lints

import (
	"testing"
)

func TestSubCertCountryNameMustAppear(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertCountryNameMustAppear.pem"
	desEnum := Error
	out, _ := Lints["e_sub_cert_country_name_must_appear"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
