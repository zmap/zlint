// lint_ext_san_registered_id_present_test.go
package lints

import (
	"testing"
)

func TestSANRegIdMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCaGood.pem"
	desEnum := Pass
	out, _ := Lints["e_ext_san_registered_id_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANRegIdPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/SANRegisteredIdBeginning.pem"
	desEnum := Error
	out, _ := Lints["e_ext_san_registered_id_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSANRegIdPresent2(t *testing.T) {
	inputPath := "../testlint/testCerts/SANRegisteredIdEnd.pem"
	desEnum := Error
	out, _ := Lints["e_ext_san_registered_id_present"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
