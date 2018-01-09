package lints

import (
	"testing"
)

func TestSubCertSkiMissing(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertNoSKI.pem"
	expected := Warn
	out := Lints["w_ext_subject_key_identifier_missing_sub_cert"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertSkiPresent(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	expected := Pass
	out := Lints["w_ext_subject_key_identifier_missing_sub_cert"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
