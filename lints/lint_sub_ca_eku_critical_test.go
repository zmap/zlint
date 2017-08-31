// lint_sub_ca_eku_critical_test.go
package lints

import (
	"testing"
)

func TestSubCaEkuCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWEkuCrit.pem"
	expected := Warn
	out := Lints["w_sub_ca_eku_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSubCaEkuNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/subCAWEkuNoCrit.pem"
	expected := Pass
	out := Lints["w_sub_ca_eku_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
