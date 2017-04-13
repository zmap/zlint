// lint_br_ian_bare_wildcard_test.go
package lints

import (
	"testing"
)

func TestBrIANBareWildcard(t *testing.T) {
	inputPath := "../testlint/testCerts/IANBareWildcard.pem"
	desEnum := Error
	out, _ := Lints["e_ian_bare_wildcard"].ExecuteTest(ReadCertificate(inputPath))
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
	out, _ := Lints["e_ian_bare_wildcard"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
