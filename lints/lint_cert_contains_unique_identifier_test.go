// lint_cert_contains_unique_identifier_test.go
package lints

import (
	"testing"
)

func TestUIDPresentIssuer(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/issuerUID.pem"
	desEnum := Error
	out, _ := Lints["e_cert_contains_unique_identifier"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestUIDPresentSubject(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/subjectUID.pem"
	desEnum := Error
	out, _ := Lints["e_cert_contains_unique_identifier"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestUIDMissing(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	desEnum := Pass
	out, _ := Lints["e_cert_contains_unique_identifier"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
