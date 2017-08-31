// lint_ext_san_missing_test.go
package lints

import (
	"testing"
)

func TestNoSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectEmptyNoSAN.pem"
	desEnum := Error
	out := Lints["e_ext_san_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestHasSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	desEnum := Pass
	out := Lints["e_ext_san_missing"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
