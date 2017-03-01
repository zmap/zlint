// lint_ext_cert_policy_explicit_text_includes_control_test.go
package lints

import (
	"testing"
)

func TestExplicitTextUtfControlX10(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/utf8ControlX10.cer"
	desEnum := Warn
	out, _ := Lints["ext_cert_policy_explicit_text_includes_control"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestExplicitTextUtfControlX88(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/utf8ControlX88.cer"
	desEnum := Warn
	out, _ := Lints["ext_cert_policy_explicit_text_includes_control"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestExplicitTextUtfNoControl(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/utf8NoControl.cer"
	desEnum := Pass
	out, _ := Lints["ext_cert_policy_explicit_text_includes_control"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
