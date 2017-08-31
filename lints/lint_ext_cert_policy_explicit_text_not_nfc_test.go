package lints

import (
	"testing"
)

func TestExplicitTextUtf8NFC(t *testing.T) {
	
	inputPath := "../testlint/testCerts/userNoticeExpTextUtf8.pem"
	expected := Pass
	out := Lints["w_ext_cert_policy_explicit_text_not_nfc"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestExplicitTextUtf8NotNFC(t *testing.T) {
	
	inputPath := "../testlint/testCerts/explicitTextUtf8NotNFC.pem"
	expected := Warn
	out := Lints["w_ext_cert_policy_explicit_text_not_nfc"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestExplicitTextBMPNFC(t *testing.T) {
	
	inputPath := "../testlint/testCerts/explicitTextBMPNFC.pem"
	expected := Pass
	out := Lints["w_ext_cert_policy_explicit_text_not_nfc"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestExplicitTextBMPNotNFC(t *testing.T) {
	
	inputPath := "../testlint/testCerts/explicitTextBMPNotNFC.pem"
	expected := Warn
	out := Lints["w_ext_cert_policy_explicit_text_not_nfc"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
