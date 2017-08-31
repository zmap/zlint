// lint_sub_cert_or_sub_ca_using_sha1_test.go
package lints

import (
	"testing"
)

// As a note, these certificates were not built, but instead grabbed from censys.io/query
// using the following query to find the raw data and match it to validity period
// select raw, parsed.validity.start from certificates.pemtificates where parsed.signature_algorithm.oid = "1.2.840.113549.1.1.5" limit 200

func TestSubCertPostalCodeProhibited(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/subCertPostalCodeProhibited.pem"
	desEnum := Error
	out := Lints["e_sub_cert_postal_code_must_not_appear"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCertPostalCodeNotProhibited(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/subCertPostalCodeNotProhibited.pem"
	desEnum := Pass
	out := Lints["e_sub_cert_postal_code_must_not_appear"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
