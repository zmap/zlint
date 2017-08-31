// lint_ext_authority_key_identifier_critical_test.go
package lints

import (
	"testing"
)

func TestAKICrit(t *testing.T) {
	inputPath := "../testlint/testCerts/akiCritical.pem"
	expected := Error
	out := Lints["e_ext_authority_key_identifier_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestAKINoCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_ext_authority_key_identifier_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


