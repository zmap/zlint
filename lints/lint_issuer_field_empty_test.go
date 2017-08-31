// lint_issuer_field_empty_test.go
package lints

import (
	"testing"
)

func TestNoIssuerField(t *testing.T) {
	inputPath := "../testlint/testCerts/issuerFieldMissing.pem"
	desEnum := Error
	out := Lints["e_issuer_field_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestHasIssuerField(t *testing.T) {
	inputPath := "../testlint/testCerts/issuerFieldFilled.pem"
	desEnum := Pass
	out := Lints["e_issuer_field_empty"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
