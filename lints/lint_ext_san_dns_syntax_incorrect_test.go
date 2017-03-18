// lint_ext_san_dns_syntax_incorrect_test.go
package lints

import (
	"testing"
)

func TestSanDNSSyntaxEndingHyphen(t *testing.T) {
	inputPath := "../testlint/testCerts/SANdnshyphensyntax.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_dns_syntax_incorrect"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanDNSSyntaxExtraPeriods(t *testing.T) {
	inputPath := "../testlint/testCerts/SANdnsbadsyntax.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_dns_syntax_incorrect"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanDNSSyntaxNotAllowedChar(t *testing.T) {
	inputPath := "../testlint/testCerts/SANdnsdollarsyntax.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_dns_syntax_incorrect"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanDNSSyntaxCorrect(t *testing.T) {
	inputPath := "../testlint/testCerts/SANdnsgoodsyntax.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_dns_syntax_incorrect"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
