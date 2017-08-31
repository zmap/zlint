// lint_cert_contains_unique_identifier_test.go
package lints

import (
	"testing"
)

func TestUIDPresentIssuer(t *testing.T) {
	
	inputPath := "../testlint/testCerts/issuerUID.pem"
	expected := Error
	out := Lints["e_cert_contains_unique_identifier"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestUIDPresentSubject(t *testing.T) {
	
	inputPath := "../testlint/testCerts/subjectUID.pem"
	expected := Error
	out := Lints["e_cert_contains_unique_identifier"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestUIDMissing(t *testing.T) {
	
	inputPath := "../testlint/testCerts/orgValGoodAllFields.pem"
	expected := Pass
	out := Lints["e_cert_contains_unique_identifier"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
