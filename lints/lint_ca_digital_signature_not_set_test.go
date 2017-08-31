// lint_ca_dig_sign_not_set_test.go
package lints

import (
	"testing"
)

func TestCaKeyUsageNoDigSign(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caKeyUsageNoCertSign.pem"
	expected := Notice
	out := Lints["n_ca_digital_signature_not_set"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestKeyUsageDigSign(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/caKeyUsageWDigSign.pem"
	expected := Pass
	out := Lints["n_ca_digital_signature_not_set"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
