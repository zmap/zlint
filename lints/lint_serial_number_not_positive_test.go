// lint_serial_number_not_positive_test.go
package lints

import (
	"testing"
)

func TestSnNeagtive(t *testing.T) {
	inputPath := "../testlint/testCerts/serialNumberNegative.pem"
	desEnum := Error
	out := Lints["e_serial_number_not_positive"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSnNotNeagtive(t *testing.T) {
	inputPath := "../testlint/testCerts/serialNumberValid.pem"
	desEnum := Pass
	out := Lints["e_serial_number_not_positive"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
