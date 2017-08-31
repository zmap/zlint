// lint_ext_subject_key_identifier_critical_test.go
package lints

import (
	"testing"
)

func TestSkiCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/skiCriticalCA.pem"
	expected := Error
	out := Lints["e_ext_subject_key_identifier_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}

func TestSkiNotCrit(t *testing.T) {
	inputPath := "../testlint/testCerts/skiNotCriticalCA.pem"
	expected := Pass
	out := Lints["e_ext_subject_key_identifier_critical"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", expected, /* The enum you expected */
			"got", out.Status, /* Actual Result */
		)
	}
}
