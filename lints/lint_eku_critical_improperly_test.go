// lint_eku_critical_improperly_test.go
package lints

import (
	"testing"
)

func TestEKUAnyCrit(t *testing.T) {
	
	inputPath := "../testlint/testCerts/ekuAnyCrit.pem"
	expected := Warn
	out := Lints["w_eku_critical_improperly"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestEKUNoCritWAny(t *testing.T) {
	
	inputPath := "../testlint/testCerts/ekuAnyNoCrit.pem"
	expected := Pass
	out := Lints["w_eku_critical_improperly"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}

func TestEKUNoAnyCrit(t *testing.T) {
	
	inputPath := "../testlint/testCerts/ekuNoAnyCrit.pem"
	expected := Pass
	out := Lints["w_eku_critical_improperly"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, 
			"expected", expected, 
			"got", out.Status, 
		)
	}
}
