package lints

import (
	"testing"
)

func TestSubCertKeyUsageNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/keyUsageNotCriticalSubCert.pem"
	expected := Warn
	out := Lints["w_ext_key_usage_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCaKeyUsageNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageNotCrit.pem"
	expected := Warn
	out := Lints["w_ext_key_usage_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertKeyUsageCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/domainValGoodSubject.pem"
	expected := Pass
	out := Lints["w_ext_key_usage_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaKeyUsageCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageCrit.pem"
	expected := Pass
	out := Lints["w_ext_key_usage_not_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertKeyUsageNotIncludedCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/caKeyUsageMissing.pem"
	expected := NA
	out := Lints["e_ext_key_usage_without_bits"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
