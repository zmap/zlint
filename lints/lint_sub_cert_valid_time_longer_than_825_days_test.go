// lint_sub_cert_valid_time_longer_than_825_days_test.go
package lints

import (
	"testing"
)

func TestSubCertValidTimeLongerThan825Days(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertOver825DaysBad.pem"
	expected := Error
	out := Lints["e_sub_cert_valid_time_longer_than_825_days"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertValidTimeLongerThan825DaysBeforeCutoff(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertOver825DaysOK.pem"
	expected := NE
	out := Lints["e_sub_cert_valid_time_longer_than_825_days"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


func TestSubCertValidTime825Days(t *testing.T) {
	inputPath := "../testlint/testCerts/subCert825DaysOK.pem"
	expected := Pass
	out := Lints["e_sub_cert_valid_time_longer_than_825_days"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
