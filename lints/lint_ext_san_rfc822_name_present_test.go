// lint_ext_san_other_name_present_test.go
package lints

import (
	"testing"
)

func TestSANEmailPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/SANRFC822Beginning.pem"
	desEnum := Error
	out := Lints["e_ext_san_rfc822_name_present"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSANEmailPresent2(t *testing.T) {
	inputPath := "../testlint/testCerts/SANRFC822End.pem"
	desEnum := Error
	out := Lints["e_ext_san_rfc822_name_present"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSANEmailMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCaGood.pem"
	desEnum := Pass
	out := Lints["e_ext_san_rfc822_name_present"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
