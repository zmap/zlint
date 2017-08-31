// lint_cert_unique_identifier_version_not_2_or_3_test.go
package lints

import (
	"testing"
)

func TestUniqueIdVersionNot1(t *testing.T) {
	inputPath := "../testlint/testCerts/uniqueIdVersion3.pem"
	desEnum := Pass
	out := Lints["e_cert_unique_identifier_version_not_2_or_3"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestUniqueIdVersion1(t *testing.T) {
	inputPath := "../testlint/testCerts/uniqueIdVersion1.pem"
	desEnum := Error
	out := Lints["e_cert_unique_identifier_version_not_2_or_3"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
