// lint_ext_san_uniform_resource_identifier_present_test.go
package lints

import (
	"testing"
)

func TestSANURIMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/SANCaGood.pem"
	expected := Pass
	out := Lints["e_ext_san_uniform_resource_identifier_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSANURIPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIBeginning.pem"
	expected := Error
	out := Lints["e_ext_san_uniform_resource_identifier_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSANURIPresent2(t *testing.T) {
	inputPath := "../testlint/testCerts/SANURIEnd.pem"
	expected := Error
	out := Lints["e_ext_san_uniform_resource_identifier_present"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
