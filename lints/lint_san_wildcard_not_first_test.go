// lint_br_san_wildcard_not_first_test.go
package lints

import (
	"testing"
)

func TestBrSANWildcardFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/SANWildcardFirst.pem"
	expected := Error
	out := Lints["e_san_wildcard_not_first"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestBrSANWildcardNotFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIValid.pem"
	expected := Pass
	out := Lints["e_san_wildcard_not_first"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
