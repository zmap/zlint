// lint_ca_dig_sign_not_set_test.go
package lints

import (
	"testing"
)

func TestCaKeyUsageNoDigSign(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caKeyUsageNoCertSign.pem"
	desEnum := Notice
	out := Lints["n_ca_digital_signature_not_set"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestKeyUsageDigSign(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caKeyUsageWDigSign.pem"
	desEnum := Pass
	out := Lints["n_ca_digital_signature_not_set"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
