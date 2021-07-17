package main

import (
	"encoding/base64"
	"encoding/pem"
	"testing"
	"time"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/util"
)

func TestRootCA(t *testing.T) {
	ca, err := newTrustAnchor()
	if err != nil {
		t.Fatal(err)
	}
	if !util.IsCACert(ca.Certificate) {
		t.Errorf("is not a ca: %s", encode(ca))
	}
	if !util.IsSelfSigned(ca.Certificate) {
		t.Errorf("is not self signed: %s", encode(ca))
	}
	if !util.IsRootCA(ca.Certificate) {
		t.Errorf("is not a root ca: %s", encode(ca))
	}
}

func TestIntermediate(t *testing.T) {
	ca, err := newTrustAnchor()
	if err != nil {
		t.Fatal(err)
	}
	intermediate, err := newIntermediate(ca)
	if err != nil {
		t.Fatal(err)
	}
	if !util.IsCACert(intermediate.Certificate) {
		t.Errorf("is not a ca: %s", encode(ca))
	}
	if util.IsSelfSigned(intermediate.Certificate) {
		t.Errorf("is self signed: %s", encode(ca))
	}
	if util.IsRootCA(intermediate.Certificate) {
		t.Errorf("is a root ca: %s", encode(ca))
	}
}

func TestLeaf(t *testing.T) {
	ca, err := newTrustAnchor()
	if err != nil {
		t.Fatal(err)
	}
	intermediate, err := newIntermediate(ca)
	if err != nil {
		t.Fatal(err)
	}
	leaf, err := newLeaf(ca, []*Certificate{intermediate})
	if err != nil {
		t.Fatal(err)
	}
	if util.IsCACert(leaf.Certificate) {
		t.Errorf("is  a ca: %s", encode(ca))
	}
	if util.IsSelfSigned(leaf.Certificate) {
		t.Errorf("is self signed: %s", encode(ca))
	}
	if util.IsRootCA(leaf.Certificate) {
		t.Errorf("is a root ca: %s", encode(ca))
	}
	if !util.IsSubscriberCert(leaf.Certificate) {
		t.Errorf("is not a subscriber: %s", encode(ca))
	}
}

func TestChainVerifies(t *testing.T) {
	ca, err := newTrustAnchor()
	if err != nil {
		t.Fatal(err)
	}
	intermediate, err := newIntermediate(ca)
	if err != nil {
		t.Fatal(err)
	}
	leaf, err := newLeaf(ca, []*Certificate{intermediate})
	if err != nil {
		t.Fatal(err)
	}
	roots := x509.NewCertPool()
	roots.AddCert(ca.Certificate)
	intermediates := x509.NewCertPool()
	intermediates.AddCert(intermediate.Certificate)
	_, _, _, err = leaf.Verify(x509.VerifyOptions{
		Intermediates: intermediates,
		Roots:         roots,
		CurrentTime:   time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestChainNoIntermediatesVerifies(t *testing.T) {
	ca, err := newTrustAnchor()
	if err != nil {
		t.Fatal(err)
	}
	leaf, err := newLeaf(ca, []*Certificate{})
	if err != nil {
		t.Fatal(err)
	}
	roots := x509.NewCertPool()
	roots.AddCert(ca.Certificate)
	_, _, _, err = leaf.Verify(x509.VerifyOptions{
		Roots:       roots,
		CurrentTime: time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestChainMultipleIntermediatesVerifies(t *testing.T) {
	ca, err := newTrustAnchor()
	if err != nil {
		t.Fatal(err)
	}
	intermediate1, err := newIntermediate(ca)
	if err != nil {
		t.Fatal(err)
	}
	intermediate2, err := newIntermediate(intermediate1)
	if err != nil {
		t.Fatal(err)
	}
	intermediate3, err := newIntermediate(intermediate2)
	if err != nil {
		t.Fatal(err)
	}
	leaf, err := newLeaf(ca, []*Certificate{intermediate1, intermediate2, intermediate3})
	if err != nil {
		t.Fatal(err)
	}
	roots := x509.NewCertPool()
	roots.AddCert(ca.Certificate)
	intermediates := x509.NewCertPool()
	intermediates.AddCert(intermediate1.Certificate)
	intermediates.AddCert(intermediate2.Certificate)
	intermediates.AddCert(intermediate3.Certificate)
	_, _, _, err = leaf.Verify(x509.VerifyOptions{
		Intermediates: intermediates,
		Roots:         roots,
		CurrentTime:   time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestBadVerify(t *testing.T) {
	badRoot := `
MIIBBTCBrKADAgECAgEBMAoGCCqGSM49BAMCMAAwIhgPMDAwMTAxMDEwMDAwMDBa
GA85OTk4MTEzMDAwMDAwMFowADBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABCJT
/KPW7GdIrQDpfeT/nSozsdWTTJvrcFSogu+qBT46SJZAzV9gVr0d1tXC52v6hsvU
QRHyQrFaFq/nzTyTBiajEzARMA8GA1UdEwEB/wQFMAMBAf8wCgYIKoZIzj0EAwID
SAAwRQIgI62LZpgjBX77r6ofW+exerSQL98gwaYri5gBNOU7+TACIQD4uZF5IGgo
wif20LYD26BzLZQTncXVx2jSzTxpQbMDgg==
`
	b, err := base64.StdEncoding.DecodeString(badRoot)
	if err != nil {
		t.Fatal(err)
	}
	badRootCert, err := x509.ParseCertificate(b)
	if err != nil {
		t.Fatal(err)
	}
	ca, err := newTrustAnchor()
	if err != nil {
		t.Fatal(err)
	}
	intermediate, err := newIntermediate(ca)
	if err != nil {
		t.Fatal(err)
	}
	leaf, err := newLeaf(ca, []*Certificate{intermediate})
	if err != nil {
		t.Fatal(err)
	}
	roots := x509.NewCertPool()
	// Setting this to the wrong root is the crux of the test.
	roots.AddCert(badRootCert)
	intermediates := x509.NewCertPool()
	intermediates.AddCert(intermediate.Certificate)
	_, _, _, err = leaf.Verify(x509.VerifyOptions{
		Intermediates: intermediates,
		Roots:         roots,
		CurrentTime:   time.Now(),
	})
	if err == nil {
		t.Fatal("generated certificate chain incorrectly verified with wrong root CA")
	}
}

func encode(c *Certificate) string {
	return string(pem.EncodeToMemory(&pem.Block{
		Type:    "CERTIFICATE",
		Headers: nil,
		Bytes:   []byte(base64.StdEncoding.EncodeToString(c.Certificate.Raw)),
	}))
}
