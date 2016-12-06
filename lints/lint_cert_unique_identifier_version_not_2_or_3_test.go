// lint_cert_unique_identifier_version_not_2_or_3_test.go
package lints

import (

	"testing"
)

func TestUniqueIdVersionNot1(t *testing.T) {
	inputPath := "../testlint/testCerts/uniqueIdVersion3.cer"
	desEnum := Pass
	out, _ := Lints["cert_unique_identifier_version_not_2_or_3"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestUniqueIdVersion1(t *testing.T) {
	inputPath := "../testlint/testCerts/uniqueIdVersion1.cer"
	desEnum := Error
	out, _ := Lints["cert_unique_identifier_version_not_2_or_3"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
