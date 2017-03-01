// lint_serial_number_longer_than_20_octets_test.go
package lints

import (
	"testing"
)

func TestSnTooLarge(t *testing.T) {
	inputPath := "../testlint/testCerts/serialNumberLarge.cer"
	desEnum := Error
	out, _ := Lints["serial_number_longer_than_20_octets"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestSnNotTooLarge(t *testing.T) {
	inputPath := "../testlint/testCerts/serialNumberValid.cer"
	desEnum := Pass
	out, _ := Lints["serial_number_longer_than_20_octets"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
