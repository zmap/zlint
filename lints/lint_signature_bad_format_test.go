// lint_signature_bad_format_test.go
package lints

import (
	"testing"
)

func TestSignatureGoodFormat(t *testing.T) {
	inputPath := "../testlint/testCerts/DSAHasParam.pem"
	desEnum := Pass
	out, _ := Lints["e_signature_bad_format"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSignatureBadFormat(t *testing.T) {
	inputPath := "../testlint/testCerts/DSANoParam.pem"
	desEnum := Error
	out, _ := Lints["e_signature_bad_format"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
