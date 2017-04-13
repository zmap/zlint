// lint_rsa_mod_not_odd_test.go
package lints

import (
	"testing"
)

func TestRsaModEven(t *testing.T) {
	inputPath := "../testlint/testCerts/evenRsaMod.pem"
	desEnum := Warn
	out, _ := Lints["w_rsa_mod_not_odd"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestRsaModOdd(t *testing.T) {
	inputPath := "../testlint/testCerts/oddRsaMod.pem"
	desEnum := Pass
	out, _ := Lints["w_rsa_mod_not_odd"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
