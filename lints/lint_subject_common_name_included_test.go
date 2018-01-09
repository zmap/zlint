package lints

import (
	"testing"
)

func TestCN(t *testing.T) {
	inputPath := "../testlint/testCerts/commonNamesURL.pem"
	expected := Notice
	out := Lints["n_subject_common_name_included"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestNoCN(t *testing.T) {
	inputPath := "../testlint/testCerts/commonNamesGood.pem"
	expected := Pass
	out := Lints["n_subject_common_name_included"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
