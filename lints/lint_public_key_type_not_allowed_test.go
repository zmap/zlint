// lint_public_key_type_not_allowed_test.go
package lints

import (
	"testing"
)

func TestPKTypeUnknown(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/unknownpublickey.pem"
	desEnum := Error
	out, _ := Lints["e_public_key_type_not_allowed"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestPKTypeRSA(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/RSAWithSHA1before2016.pem"
	desEnum := Pass
	out, _ := Lints["e_public_key_type_not_allowed"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestPKTypeECDSA(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/ecdsaP256.pem"
	desEnum := Pass
	out, _ := Lints["e_public_key_type_not_allowed"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
