// lint_eku_critical_improperly_test.go
package lints

import (
	"testing"
)

func TestEKUAnyCrit(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/ekuAnyCrit.pem"
	expected := Warn
	out := Lints["w_eku_critical_improperly"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestEKUNoCritWAny(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/ekuAnyNoCrit.pem"
	expected := Pass
	out := Lints["w_eku_critical_improperly"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestEKUNoAnyCrit(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/ekuNoAnyCrit.pem"
	expected := Pass
	out := Lints["w_eku_critical_improperly"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
