// lint_br_ian_wildcard_not_first_test.go
package lints

import (
	"testing"
)

func TestBrIANWildcardFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/IANWildcardFirst.pem"
	expected := Error
	out := Lints["e_ian_wildcard_not_first"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestBrIANWildcardNotFirst(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.pem"
	expected := Pass
	out := Lints["e_ian_wildcard_not_first"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
