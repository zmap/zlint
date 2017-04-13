// lint_ext_san_dns_syntax_incorrect_test.go
package lints

import (
	"testing"
)

func TestSANDNSSyntaxEndingHyphen(t *testing.T) {
	inputPath := "../testlint/testCerts/SANdnshyphensyntax.pem"
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

func TestSANDNSSyntaxExtraPeriods(t *testing.T) {
	inputPath := "../testlint/testCerts/SANdnsbadsyntax.pem"
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

func TestSANDNSSyntaxNotAllowedChar(t *testing.T) {
	inputPath := "../testlint/testCerts/SANdnsdollarsyntax.pem"
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

func TestSANDNSSyntaxCorrect(t *testing.T) {
	inputPath := "../testlint/testCerts/SANdnsgoodsyntax.pem"
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
