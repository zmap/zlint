// lint_wrong_time_format_pre2050_test.go
package lints

import (
	"testing"
)

func TestGeneralizedAfter2050(t *testing.T) {
	inputPath := "../testlint/testCerts/generalizedAfter2050.pem"
	expected := Pass
	out := Lints["e_wrong_time_format_pre2050"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestUTCPrior2050(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_wrong_time_format_pre2050"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestGeneralizedPrior2050(t *testing.T) {
	inputPath := "../testlint/testCerts/generalizedPrior2050.pem"
	expected := Error
	out := Lints["e_wrong_time_format_pre2050"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
