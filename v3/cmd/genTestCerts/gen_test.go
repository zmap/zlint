package main

import (
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"strings"
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
	current, expired, never, err := leaf.Verify(x509.VerifyOptions{
		Intermediates: intermediates,
		Roots:         roots,
		CurrentTime:   time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	assertChains(current, expired, never, 1, t)
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
	current, expired, never, err := leaf.Verify(x509.VerifyOptions{
		Roots:       roots,
		CurrentTime: time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	assertChains(current, expired, never, 1, t)
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
	current, expired, never, err := leaf.Verify(x509.VerifyOptions{
		Intermediates: intermediates,
		Roots:         roots,
		CurrentTime:   time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	assertChains(current, expired, never, 1, t)
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
	current, expired, never, err := leaf.Verify(x509.VerifyOptions{
		Intermediates: intermediates,
		Roots:         roots,
		CurrentTime:   time.Now(),
	})
	if err == nil {
		t.Fatal("generated certificate chain incorrectly verified with wrong root CA")
	}
	assertChains(current, expired, never, 0, t)
}

func TestGetTestData(t *testing.T) {
	got, err := getTestDataDir()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasSuffix(got, "zlint/v3/testdata") {
		t.Fatalf("wanted path ending in 'zlint/v3/testdata' got '%s'", got)
	}
}

func TestSaveCert(t *testing.T) {
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
	fname, err := saveCertificateToTestdata(leaf, "UNIT_TEST.pem")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(fname)
	_, err = os.Stat(fname)
	if err != nil {
		t.Fatal(err)
	}
}

func assertChains(current, expired, never []x509.CertificateChain, currentWant int, t *testing.T) {
	expiredWant := 0
	neverWant := 0
	if len(current) != currentWant {
		b := strings.Builder{}

		b.WriteString(fmt.Sprintf("got %d valid certificate chains, wanted %d\n", len(current), currentWant))
		for i, chain := range current {
			b.WriteString(fmt.Sprintf("chain #%d\n", i+1))
			b.WriteString(encodeChain(chain))
		}
		t.Error(b.String())
	}
	if len(expired) != expiredWant {
		b := strings.Builder{}
		b.WriteString(fmt.Sprintf("got %d expired certificate chains, wanted %d\n", len(expired), expiredWant))
		for i, chain := range expired {
			b.WriteString(fmt.Sprintf("chain #%d\n", i+1))
			b.WriteString(encodeChain(chain))
		}
		t.Error(b.String())
	}
	if len(never) != neverWant {
		b := strings.Builder{}
		b.WriteString(fmt.Sprintf("got %d 'never' certificate chains, wanted %d\n", len(never), neverWant))
		for i, chain := range never {
			b.WriteString(fmt.Sprintf("chain #%d\n", i+1))
			b.WriteString(encodeChain(chain))
		}
		t.Error(b.String())
	}
}

func encode(c *Certificate) string {
	return encodeX509(c.Certificate)
}

func encodeX509(c *x509.Certificate) string {
	return string(pem.EncodeToMemory(&pem.Block{
		Type:    "CERTIFICATE",
		Headers: nil,
		Bytes:   []byte(base64.StdEncoding.EncodeToString(c.Raw)),
	}))
}

func encodeChain(chain x509.CertificateChain) string {
	b := strings.Builder{}
	for _, cert := range chain {
		s, err := openSSLFormatCertificate(&Certificate{Certificate: cert})
		if err != nil {
			panic(err)
		}
		b.WriteString(s)
		b.WriteByte('\n')
	}
	return b.String()
}
