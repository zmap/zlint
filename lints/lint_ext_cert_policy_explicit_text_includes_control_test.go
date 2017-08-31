// lint_ext_cert_policy_explicit_text_includes_control_test.go
package lints

import (
	"testing"
)

func TestExplicitTextUtfControlX10(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/utf8ControlX10.pem"
	expected := Warn
	out := Lints["w_ext_cert_policy_explicit_text_includes_control"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestExplicitTextUtfControlX88(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/utf8ControlX88.pem"
	expected := Warn
	out := Lints["w_ext_cert_policy_explicit_text_includes_control"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestExplicitTextUtfNoControl(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/utf8NoControl.pem"
	expected := Pass
	out := Lints["w_ext_cert_policy_explicit_text_includes_control"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
