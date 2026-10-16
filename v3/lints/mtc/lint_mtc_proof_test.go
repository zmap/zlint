package mtc

import (
	"encoding/pem"
	"os"
	"testing"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
	"golang.org/x/crypto/cryptobyte"
	cbasn1 "golang.org/x/crypto/cryptobyte/asn1"
)

func TestMTCProof(t *testing.T) {
	cert := readMTCProof(t, "mtc_plants07_valid.pem")
	for _, name := range []string{"e_mtc_proof_structure_valid", "e_mtc_proof_self_contained_valid"} {
		t.Run(name, func(t *testing.T) {
			result := test.TestLintCert(name, cert, lint.NewEmptyConfig())
			if result.Status != lint.Pass {
				t.Fatalf("got %s: %s", result.Status, result.Details)
			}
		})
	}
}

func readMTCProof(t *testing.T, filename string) *x509.Certificate {
	t.Helper()
	data, err := os.ReadFile("../../testdata/" + filename)
	if err != nil {
		t.Fatal(err)
	}
	var block *pem.Block
	for len(data) > 0 {
		block, data = pem.Decode(data)
		if block == nil {
			t.Fatal("failed to decode test certificate")
		}
		if block.Type == "CERTIFICATE" {
			break
		}
	}
	if block == nil || block.Type != "CERTIFICATE" {
		t.Fatal("certificate block not found")
	}
	input := cryptobyte.String(block.Bytes)
	var certificate, tbs cryptobyte.String
	if !input.ReadASN1(&certificate, cbasn1.SEQUENCE) {
		t.Fatal("failed to extract certificate")
	}
	if !certificate.ReadASN1Element(&tbs, cbasn1.SEQUENCE) {
		t.Fatal("failed to extract TBSCertificate")
	}
	return &x509.Certificate{Raw: block.Bytes, RawTBSCertificate: tbs}
}
