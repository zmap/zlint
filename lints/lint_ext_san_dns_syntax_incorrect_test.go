// lint_ext_san_dns_syntax_incorrect_test.go
package lints

import (
	"testing"
)

func TestSanDNSSyntaxEndingHyphen(t *testing.T) {
	inputPath := "../testlint/testCerts/sandnshyphensyntax.cer"
	desEnum := Error
	out, _ := Lints["ext_san_dns_syntax_incorrect"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanDNSSyntaxExtraPeriods(t *testing.T) {
	inputPath := "../testlint/testCerts/sandnsbadsyntax.cer"
	desEnum := Error
	out, _ := Lints["ext_san_dns_syntax_incorrect"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanDNSSyntaxNotAllowedChar(t *testing.T) {
	inputPath := "../testlint/testCerts/sandnsdollarsyntax.cer"
	desEnum := Error
	out, _ := Lints["ext_san_dns_syntax_incorrect"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanDNSSyntaxCorrect(t *testing.T) {
	inputPath := "../testlint/testCerts/sandnsgoodsyntax.cer"
	desEnum := Pass
	out, _ := Lints["ext_san_dns_syntax_incorrect"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
