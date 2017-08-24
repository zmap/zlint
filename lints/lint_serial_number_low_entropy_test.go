// lint_serial_number_not_positive_test.go
package lints

import (
	"testing"
)

func TestSnLowEntropy(t *testing.T) {
	inputPath := "../testlint/testCerts/serialNumberLowEntropy.pem"
	desEnum := Warn
	out, _ := Lints["e_serial_number_low_entropy"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
