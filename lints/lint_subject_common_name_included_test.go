// lint_subject_common_name_included_test.go
package lints

import (
	"testing"
)

func TestCN(t *testing.T) {
	inputPath := "../testlint/testCerts/commonNamesURL.pem"
	expected := Notice
	out := Lints["n_subject_common_name_included"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestNoCN(t *testing.T) {
	inputPath := "../testlint/testCerts/commonNamesGood.pem"
	expected := Pass
	out := Lints["n_subject_common_name_included"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
