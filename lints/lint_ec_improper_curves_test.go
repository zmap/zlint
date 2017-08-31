// lint_ec_improper_curves_test.go
package lints

import (
	"testing"
)

func TestECP224(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/ecdsaP224.pem"
	expected := Error
	out := Lints["e_ec_improper_curves"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestECP256(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/ecdsaP256.pem"
	expected := Pass
	out := Lints["e_ec_improper_curves"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestECP384(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/ecdsaP384.pem"
	expected := Pass
	out := Lints["e_ec_improper_curves"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestECP521(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/ecdsaP521.pem"
	expected := Pass
	out := Lints["e_ec_improper_curves"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
