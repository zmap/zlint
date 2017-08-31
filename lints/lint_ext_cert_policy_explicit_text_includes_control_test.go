// lint_ext_cert_policy_explicit_text_includes_control_test.go
package lints

import (
	"testing"
)

func TestExplicitTextUtfControlX10(t *testing.T) {
	inputPath := "../testlint/testCerts/utf8ControlX10.pem"
	expected := Warn
	out := Lints["w_ext_cert_policy_explicit_text_includes_control"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExplicitTextUtfControlX88(t *testing.T) {
	inputPath := "../testlint/testCerts/utf8ControlX88.pem"
	expected := Warn
	out := Lints["w_ext_cert_policy_explicit_text_includes_control"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestExplicitTextUtfNoControl(t *testing.T) {
	inputPath := "../testlint/testCerts/utf8NoControl.pem"
	expected := Pass
	out := Lints["w_ext_cert_policy_explicit_text_includes_control"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


