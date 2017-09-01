// lint_utc_time_not_in_zulu_test.go
package lints

import (
	"testing"
)

func TestUtcZulu(t *testing.T) {
	inputPath := "../testlint/testCerts/utcHasSeconds.pem"
	expected := Pass
	out := Lints["e_utc_time_not_in_zulu"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestUtcNotZulu(t *testing.T) {
	inputPath := "../testlint/testCerts/utcNotZulu.pem"
	expected := Error
	out := Lints["e_utc_time_not_in_zulu"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
