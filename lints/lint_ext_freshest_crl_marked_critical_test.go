package lints

import (
	"testing"
)

func TestFreshestCrlCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/frshCRLCritical.pem"
	expected := Error
	out := Lints["e_ext_freshest_crl_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestFreshestCrlNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/frshCRLNotCritical.pem"
	expected := Pass
	out := Lints["e_ext_freshest_crl_marked_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
