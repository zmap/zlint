// lint_serial_number_longer_than_20_octets_test.go
package lints

import (
	"testing"
)

func TestSnTooLarge(t *testing.T) {
	inputPath := "../testlint/testCerts/serialNumberLarge.pem"
	expected := Error
	out := Lints["e_serial_number_longer_than_20_octets"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSnNotTooLarge(t *testing.T) {
	inputPath := "../testlint/testCerts/serialNumberValid.pem"
	expected := Pass
	out := Lints["e_serial_number_longer_than_20_octets"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
