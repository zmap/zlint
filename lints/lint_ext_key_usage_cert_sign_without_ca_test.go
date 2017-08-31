// lint_ext_key_usage_cert_sign_without_ca_test.go
package lints

import (
	"testing"
)

func TestCertSignNoCa(t *testing.T) {
	inputPath := "../testlint/testCerts/keyUsageCertSignNoBC.pem"
	desEnum := Error
	out := Lints["e_ext_key_usage_cert_sign_without_ca"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCertSignIsCa(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageNoCertSign.pem"
	desEnum := Pass
	out := Lints["e_ext_key_usage_cert_sign_without_ca"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
