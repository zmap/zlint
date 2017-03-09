// lint_ext_san_uniform_resource_identifier_present_test.go
package lints

import (
	"testing"
)

func TestSanURIMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/sanCaGood.cer"
	desEnum := Pass
	out, _ := Lints["e_ext_san_uniform_resource_identifier_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanURIPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/sanURIBeginning.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_uniform_resource_identifier_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSanURIPresent2(t *testing.T) {
	inputPath := "../testlint/testCerts/sanURIEnd.cer"
	desEnum := Error
	out, _ := Lints["e_ext_san_uniform_resource_identifier_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
