/*
 * ZLint Copyright 2024 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package rfc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	"testing"
	"time"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

const lintUnderTest = "e_crl_revoked_certificates_field_must_be_empty"

// generateCRLFromTemplate takes a CRL template and creates an in-memory issuer
// capable of signing the CRL and returns the resulting CRL or an error.
func generateCRLFromTemplate(crlTemplate *x509.RevocationList) (*x509.RevocationList, error) {
	signer, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	issuerTemplate := &x509.Certificate{
		SerialNumber:          big.NewInt(666),
		BasicConstraintsValid: true,
		IsCA:                  true,
		Subject: pkix.Name{
			CommonName: "Big CA",
		},
		SubjectKeyId: []byte{1, 2, 3, 4, 5, 6},
		KeyUsage:     x509.KeyUsageCRLSign | x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
	}

	issuerBytes, err := x509.CreateCertificate(rand.Reader, issuerTemplate, issuerTemplate, signer.Public(), signer)
	if err != nil {
		return nil, err
	}

	issuer, err := x509.ParseCertificate(issuerBytes)
	if err != nil {
		return nil, err
	}
	crlBytes, err := x509.CreateRevocationList(rand.Reader, crlTemplate, issuer, signer)
	if err != nil {
		return nil, err
	}

	// We're not going to trust that x509.ParseRevocationList is doing the
	// correct thing here and will instead parse the DER-encoded ASN.1 bytes of
	// this CRL later on in the lint itself.
	actualCRL, err := x509.ParseRevocationList(crlBytes)
	if err != nil {
		return nil, err
	}

	return actualCRL, nil
}

var defaultTemplate = x509.RevocationList{
	Number:     big.NewInt(1),
	ThisUpdate: time.Now().Add(-time.Second),
	NextUpdate: time.Now().Add(10 * time.Second),
}

func TestEmptyRevokedCertificatesField(t *testing.T) {
	crlTemplate := defaultTemplate
	crlTemplate.RevokedCertificates = []x509.RevokedCertificate{}

	crl, err := generateCRLFromTemplate(&crlTemplate)
	if err != nil {
		t.Error(err)
	}

	out := test.TestLintRevocationList(t, lintUnderTest, crl, lint.NewEmptyConfig())
	expected := lint.Pass
	if out.Status != expected {
		t.Errorf("expected %s, got %s", expected, out.Status)
	}

	expectedRevokedCerts := 0
	if len(crl.RevokedCertificates) != expectedRevokedCerts {
		t.Errorf("expected %d revoked certificates in CRL, got %d", expectedRevokedCerts, len(crl.RevokedCertificates))
	}
}

func TestNilRevokedCertificatesField(t *testing.T) {
	crlTemplate := defaultTemplate
	crlTemplate.RevokedCertificates = nil

	crl, err := generateCRLFromTemplate(&crlTemplate)
	if err != nil {
		t.Error(err)
	}

	out := test.TestLintRevocationList(t, lintUnderTest, crl, lint.NewEmptyConfig())
	expected := lint.Pass
	if out.Status != expected {
		t.Errorf("expected %s, got %s", expected, out.Status)
	}

	expectedRevokedCerts := 0
	if len(crl.RevokedCertificates) != expectedRevokedCerts {
		t.Errorf("expected %d revoked certificates in CRL, got %d", expectedRevokedCerts, len(crl.RevokedCertificates))
	}
}

func TestPopulatedRevokedCertificatesField(t *testing.T) {
	crlTemplate := defaultTemplate
	crlTemplate.RevokedCertificates = append(crlTemplate.RevokedCertificates, x509.RevokedCertificate{
		SerialNumber:   big.NewInt(876),
		RevocationTime: time.Now().Add(-24 * time.Hour),
	})

	crl, err := generateCRLFromTemplate(&crlTemplate)
	if err != nil {
		t.Error(err)
	}

	// The lint should not run for this case because we populated the
	// TBSCertList.revokedCertificates field.
	expected := lint.NA
	out := test.TestLintRevocationList(t, lintUnderTest, crl, lint.NewEmptyConfig())
	if out.Status != expected {
		t.Errorf("expected %s, got %s", expected, out.Status)
	}

	expectedRevokedCerts := 1
	if len(crl.RevokedCertificates) != expectedRevokedCerts {
		t.Errorf("expected %d revoked certificates in CRL, got %d", expectedRevokedCerts, len(crl.RevokedCertificates))
	}
}

func TestRevokedCertificatesContainerExistsButIsEmpty(t *testing.T) {
	// Negative test data created outside the purview of Golang.
	badCRLFiles := []string{
		"crlWithRevokedCertificatesContainerButNoActualRevokedCerts-ReallyReallyBroken.pem",
		"crlWithRevokedCertificatesContainerButNoActualRevokedCerts-CBonnell.pem",
		"crlEntrustNoRevokedCerts01.pem", // https://bugzilla.mozilla.org/show_bug.cgi?id=1889217
		"crlEntrustNoRevokedCerts02.pem",
	}

	expected := lint.Error
	for _, crl := range badCRLFiles {
		out := test.TestRevocationListLint(t, lintUnderTest, crl)
		if out.Status != expected {
			t.Errorf("expected %s, got %s", expected, out.Status)
		}
	}
}
