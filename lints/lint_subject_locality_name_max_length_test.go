// lint_subject_locality_name_max_length_test.go
package lints

import (
	"testing"
)

func TestSubjectLocalityNameLengthGood(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectLocalityNameLengthGood.pem"
	expected := Pass
	out := Lints["e_subject_locality_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestSubjectLocalityNameLong(t *testing.T) {
	inputPath := "../testlint/testCerts/subjectLocalityNameLong.pem"
	expected := Error
	out := Lints["e_subject_locality_name_max_length"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
