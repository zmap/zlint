package lints

import (
	"testing"
)

func crlCompleteDp(t *testing.T) {
	inputPath := "../testlint/testCerts/crlComlepteDp.pem"
	expected := Pass
	out := Lints["e_distribution_point_incomplete"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func crlIncompleteDp(t *testing.T) {
	inputPath := "../testlint/testCerts/crlIncomlepteDp.pem"
	expected := Error
	out := Lints["e_distribution_point_incomplete"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
