package lints

import (
	"testing"
)

func TestEvValidTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/evValidTooLong.pem"
	expected := Error
	out := Lints["e_ev_valid_time_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestEvValidNotTooLong(t *testing.T) {
	inputPath := "../testlint/testCerts/evValidNotTooLong.pem"
	expected := Pass
	out := Lints["e_ev_valid_time_too_long"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
