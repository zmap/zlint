package lints

import (
	"testing"
)

func TestEvNoBiz(t *testing.T) {
	inputPath := "../testlint/testCerts/evAllGood.pem"
	expected := Error
	out := Lints["e_ev_business_category_missing"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}


