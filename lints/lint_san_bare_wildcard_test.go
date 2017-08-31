// lint_br_san_bare_wildcard_test.go
package lints

import (
	"testing"
)

func TestBrSANBareWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/SANBareWildcard.pem"
	desEnum := Error
	out := Lints["e_san_bare_wildcard"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestBrSANNotBareWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIValid.pem"
	desEnum := Pass
	out := Lints["e_san_bare_wildcard"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
