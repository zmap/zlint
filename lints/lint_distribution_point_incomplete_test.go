// lint_distribution_point_incomplete_test.go
package lints

import (
	"testing"
)

func crlCompleteDp(t *testing.T) {
	
	inputPath := "../testlint/testCerts/crlComlepteDp.pem"
	expected := Pass
	out := Lints["e_distribution_point_incomplete"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func crlIncompleteDp(t *testing.T) {
	
	inputPath := "../testlint/testCerts/crlIncomlepteDp.pem"
	expected := Error
	out := Lints["e_distribution_point_incomplete"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
