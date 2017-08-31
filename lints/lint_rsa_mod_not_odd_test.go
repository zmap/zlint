// lint_rsa_mod_not_odd_test.go
package lints

import (
	"testing"
)

func TestRsaModEven(t *testing.T) {
	inputPath := "../testlint/testCerts/evenRsaMod.pem"
	desEnum := Warn
	out := Lints["w_rsa_mod_not_odd"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestRsaModOdd(t *testing.T) {
	inputPath := "../testlint/testCerts/oddRsaMod.pem"
	desEnum := Pass
	out := Lints["w_rsa_mod_not_odd"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
