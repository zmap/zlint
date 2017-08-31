// lint_ext_san_no_entries_test.go
package lints

import (
	"testing"
)

func TestSANNoEntry(t *testing.T) {
	inputPath := "../testlint/testCerts/SANNoEntries.pem"
	expected := Error
	out := Lints["e_ext_san_no_entries"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSANHasEntry(t *testing.T) {
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_ext_san_no_entries"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


