// lint_serial_number_not_positive_test.go
package lints

import (
	"testing"
)

func TestSnLowEntropy(t *testing.T) {
	inputPath := "../testlint/testCerts/serialNumberLowEntropy.pem"
	expected := Warn
	out := Lints["w_serial_number_low_entropy"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
