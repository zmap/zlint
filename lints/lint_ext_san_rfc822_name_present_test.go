// lint_ext_san_other_name_present_test.go
package lints

import (
	"testing"
)

func TestSanEmailPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/SANRFC822Beginning.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_rfc822_name_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanEmailPresent2(t *testing.T) {
	inputPath := "../testlint/testCerts/SANRFC822End.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_rfc822_name_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanEmailMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCaGood.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_rfc822_name_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
