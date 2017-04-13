// lint_ext_authority_key_identifier_no_key_identifier_test.go
package lints

import (
	"testing"
)

func TestAKINoKeyID(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/akiNoKeyIdentifier.pem"
	desEnum := Error
	out, _ := Lints["e_ext_authority_key_identifier_no_key_identifier"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestAKIKeyID(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	desEnum := Pass
	out, _ := Lints["e_ext_authority_key_identifier_no_key_identifier"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestAKINoKeyIDOnRoot(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/rootCANoKeyIdentifiers.pem"
	desEnum := Pass
	out, _ := Lints["e_ext_authority_key_identifier_no_key_identifier"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
