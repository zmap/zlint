// lint_cert_contains_unique_identifier_test.go
package lints

import (
	"testing"
)

func TestUIDPresentIssuer(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/issuerUID.pem"
	desEnum := Error
	out := Lints["e_cert_contains_unique_identifier"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestUIDPresentSubject(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/subjectUID.pem"
	desEnum := Error
	out := Lints["e_cert_contains_unique_identifier"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestUIDMissing(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	desEnum := Pass
	out := Lints["e_cert_contains_unique_identifier"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
