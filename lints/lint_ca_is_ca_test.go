// lint_ca_key_cert_sign_not_set_test.go
package lints

import (
	"testing"
)

func TestKeyCertSignNotCA(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/keyCertSignNotCA.pem"
	desEnum := Error
	out := Lints["e_ca_is_ca"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestKeyCertSignCA(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/keyCertSignCA.pem"
	desEnum := Pass
	out := Lints["e_ca_is_ca"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
