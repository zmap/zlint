package lints

import (
	"testing"
)

func TestSkiCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/skiCriticalCA.pem"
	expected := Error
	out := Lints["e_ext_subject_key_identifier_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSkiNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/skiNotCriticalCA.pem"
	expected := Pass
	out := Lints["e_ext_subject_key_identifier_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
