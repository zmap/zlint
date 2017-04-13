// lint_serial_number_not_positive_test.go
package lints

import (
	"testing"
)

func TestSnNeagtive(t *testing.T) {
	inputPath := "../testlint/testCerts/serialNumberNegative.pem"
	desEnum := Error
	out, _ := Lints["e_serial_number_not_positive"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSnNotNeagtive(t *testing.T) {
	inputPath := "../testlint/testCerts/serialNumberValid.pem"
	desEnum := Pass
	out, _ := Lints["e_serial_number_not_positive"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
