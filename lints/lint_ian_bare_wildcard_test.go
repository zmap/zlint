// lint_br_ian_bare_wildcard_test.go
package lints

import (
	"testing"
)

func TestBrIANBareWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/IANBareWildcard.pem"
	desEnum := Error
	out := Lints["e_ian_bare_wildcard"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBrIANNotBareWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/IANURIValid.pem"
	desEnum := Pass
	out := Lints["e_ian_bare_wildcard"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
