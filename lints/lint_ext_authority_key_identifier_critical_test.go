// lint_ext_authority_key_identifier_critical_test.go
package lints

import (
	"testing"
)

func TestAKICrit(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/akiCritical.pem"
	desEnum := Error
	out := Lints["e_ext_authority_key_identifier_critical"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestAKINoCrit(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	desEnum := Pass
	out := Lints["e_ext_authority_key_identifier_critical"].Execute(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
