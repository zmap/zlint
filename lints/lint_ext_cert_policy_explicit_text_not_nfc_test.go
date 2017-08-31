package lints

import (
	"testing"
)

func TestExplicitTextUtf8NFC(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/userNoticeExpTextUtf8.pem"
	expected := Pass
	out := Lints["w_ext_cert_policy_explicit_text_not_nfc"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestExplicitTextUtf8NotNFC(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/explicitTextUtf8NotNFC.pem"
	expected := Warn
	out := Lints["w_ext_cert_policy_explicit_text_not_nfc"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestExplicitTextBMPNFC(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/explicitTextBMPNFC.pem"
	expected := Pass
	out := Lints["w_ext_cert_policy_explicit_text_not_nfc"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestExplicitTextBMPNotNFC(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/explicitTextBMPNotNFC.pem"
	expected := Warn
	out := Lints["w_ext_cert_policy_explicit_text_not_nfc"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
