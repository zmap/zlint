// lint_ext_ian_critical_test.go
package lints

import (
	"testing"
)

func TestIANCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/IANCritical.pem"
	desEnum := Warn
	out, _ := Lints["w_ext_ian_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestIANNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/IANNotCritical.pem"
	desEnum := Pass
	out, _ := Lints["w_ext_ian_critical"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
