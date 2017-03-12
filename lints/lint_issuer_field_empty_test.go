// lint_issuer_field_empty_test.go
package lints

import (
	"testing"
)

func TestNoIssuerField(t *testing.T) {
	inputPath := "../testlint/testCerts/issuerFieldMissing.cer"
	desEnum := Error
	out, _ := Lints["e_issuer_field_empty"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestHasIssuerField(t *testing.T) {
	inputPath := "../testlint/testCerts/issuerFieldFilled.cer"
	desEnum := Pass
	out, _ := Lints["e_issuer_field_empty"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
