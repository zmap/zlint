// lint_sub_cert_or_sub_ca_using_sha1_test.go
package lints

import (
	"testing"
)

// As a note, these certificates were not built, but instead grabbed from censys.io/query
// using the following query to find the raw data and match it to validity period
// select raw, parsed.validity.start from certificates.pemtificates where parsed.signature_algorithm.oid = "1.2.840.113549.1.1.5" limit 200

func TestSHA1After2016(t *testing.T) {
	inputPath := "../testlint/testCerts/rsawithsha1after2016.pem"
	expected := Error
	out := Lints["e_sub_cert_or_sub_ca_using_sha1"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestSHA1Before2016(t *testing.T) {
	inputPath := "../testlint/testCerts/rsawithsha1before2016.pem"
	expected := NE
	out := Lints["e_sub_cert_or_sub_ca_using_sha1"].Execute(ReadCertificate(inputPath))
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}
