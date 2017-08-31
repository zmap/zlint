// lint_ext_ian_critical_test.go
package lints

import (
	"testing"
)

func TestIANCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/IANCritical.pem"
	desEnum := Warn
	out := Lints["w_ext_ian_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestIANNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/IANNotCritical.pem"
	desEnum := Pass
	out := Lints["w_ext_ian_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
