// lint_eku_critical_improperly_test.go
package lints

import (

	"testing"
)

func TestEKUAnyCrit(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/ekuAnyCrit.cer"
	desEnum := Warn
	out, _ := Lints["eku_critical_improperly"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestEKUNoCritWAny(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/ekuAnyNoCrit.cer"
	desEnum := Pass
	out, _ := Lints["eku_critical_improperly"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestEKUNoAnyCrit(t *testing.T) {
	// Only need to change these two values and the lint name
	inputPath := "../testlint/testCerts/ekuNoAnyCrit.cer"
	desEnum := Pass
	out, _ := Lints["eku_critical_improperly"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
