// lint_cert_contains_unique_identifier_test.go
package lints

import (
	"testing"
)

func TestUIDPresentIssuer(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/issuerUID.pem"
	expected := Error
	out := Lints["e_cert_contains_unique_identifier"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestUIDPresentSubject(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/subjectUID.pem"
	expected := Error
	out := Lints["e_cert_contains_unique_identifier"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestUIDMissing(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_cert_contains_unique_identifier"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
