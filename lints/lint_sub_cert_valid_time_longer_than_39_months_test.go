// lint_sub_cert_valid_time_longer_than_39_months_test.go
package lints

import (
	"testing"
)

func TestSubCertValidTimeLongerThan39Months(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertValidTimeTooLong.pem"
	expected := Error
	out := Lints["e_sub_cert_valid_time_longer_than_39_months"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertValidTimeGood(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertValidTimeGood.pem"
	expected := Pass
	out := Lints["e_sub_cert_valid_time_longer_than_39_months"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
