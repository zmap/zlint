package lints

import (
	"testing"
)

func TestExplicitTextUtf8NFC(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/userNoticeExpTextUtf8.pem"
	desEnum := Pass
	out, _ := Lints["w_ext_cert_policy_explicit_text_not_nfc"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestExplicitTextUtf8NotNFC(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/explicitTextUtf8NotNFC.pem"
	desEnum := Warn
	out, _ := Lints["w_ext_cert_policy_explicit_text_not_nfc"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestExplicitTextBMPNFC(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/explicitTextBMPNFC.pem"
	desEnum := Pass
	out, _ := Lints["w_ext_cert_policy_explicit_text_not_nfc"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestExplicitTextBMPNotNFC(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/explicitTextBMPNotNFC.pem"
	desEnum := Warn
	out, _ := Lints["w_ext_cert_policy_explicit_text_not_nfc"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
