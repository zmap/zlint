// lint_public_key_type_not_allowed_test.go
package lints

import (
	"testing"
)

func TestPKTypeUnknown(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/unknownpublickey.pem"
	desEnum := Error
	out := Lints["e_public_key_type_not_allowed"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestPKTypeRSA(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rsawithsha1before2016.pem"
	desEnum := Pass
	out := Lints["e_public_key_type_not_allowed"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestPKTypeECDSA(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/ecdsaP256.pem"
	desEnum := Pass
	out := Lints["e_public_key_type_not_allowed"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
