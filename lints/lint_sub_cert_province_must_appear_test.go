package lints

import (
	"testing"
)

// As a note, these certificates were not built, but instead grabbed from censys.io/query
// using the following query to find the raw data and match it to validity period
// select raw, parsed.validity.start from certificates.pemtificates where parsed.signature_algorithm.oid = "1.2.840.113549.1.1.5" limit 200

func TestSubCertProvinceProhibited(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertProvinceProhibited.pem"
	expected := Error
	out := Lints["e_sub_cert_province_must_appear"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSubCertProvinceNotProhibited(t *testing.T) {
	inputPath := "../testlint/testCerts/subCertProvinceNotProhibited.pem"
	expected := Pass
	out := Lints["e_sub_cert_province_must_appear"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
