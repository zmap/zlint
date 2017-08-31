// lint_subject_country_not_iso_test.go
package lints

import (
	"testing"
)

func TestGenralizedNotZulu(t *testing.T) {
	inputPath := "../testlint/testCerts/generalizedNotZulu.pem"
	expected := Error
	out := Lints["e_generalized_time_not_in_zulu"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestGenralizedZulu(t *testing.T) {
	inputPath := "../testlint/testCerts/generalizedHasSeconds.pem"
	expected := Pass
	out := Lints["e_generalized_time_not_in_zulu"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


