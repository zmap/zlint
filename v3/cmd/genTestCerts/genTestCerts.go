package main

/*
 * ZLint Copyright 2021 Regents of the University of Michigan
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

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/pem"
	"fmt"
	"math/big"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
)

// Generates a CA, an intermediate, and a leaf certificate and prints their
// OpenSSL textual output to stdout.
func main() {
	ca, err := newTrustAnchor()
	if err != nil {
		panic(err)
	}
	printCertificate(ca, "Trust Anchor")
	intermediate, err := newIntermediate(ca)
	if err != nil {
		panic(err)
	}
	printCertificate(intermediate, "Intermediate")
	leaf, err := newLeaf(ca, []*Certificate{intermediate})
	if err != nil {
		panic(err)
	}
	printCertificate(leaf, "Leaf")
}

func printCertificate(certificate *Certificate, header string) {
	fmted, err := openSSLFormatCertificate(certificate)
	if err != nil {
		panic(err)
	}
	fmt.Printf("-------------%s-------------\n", header)
	fmt.Println(fmted)
}

// This is NOT a healthy example of a CA certificate, this is nothing
// more than a self signed certificate with IsCA set to true. Not even any
// basic constraints are defined. Please do not think that this will be
// acceptable to any system, let alone lint particularly well.
func newTrustAnchor() (*Certificate, error) {
	// Edit this template to look like whatever trust anchor you need.
	template := x509.Certificate{
		Raw:                         nil,
		RawTBSCertificate:           nil,
		RawSubjectPublicKeyInfo:     nil,
		RawSubject:                  nil,
		RawIssuer:                   nil,
		Signature:                   nil,
		SignatureAlgorithm:          0,
		PublicKeyAlgorithm:          0,
		PublicKey:                   nil,
		Version:                     0,
		SerialNumber:                nextSerial(),
		Issuer:                      pkix.Name{},
		Subject:                     pkix.Name{},
		NotBefore:                   time.Time{},
		NotAfter:                    time.Date(9999, 0, 0, 0, 0, 0, 0, time.UTC),
		KeyUsage:                    0,
		Extensions:                  nil,
		ExtraExtensions:             nil,
		UnhandledCriticalExtensions: nil,
		ExtKeyUsage:                 nil,
		UnknownExtKeyUsage:          nil,
		BasicConstraintsValid:       true,
		IsCA:                        true,
		MaxPathLen:                  0,
		MaxPathLenZero:              false,
		SubjectKeyId:                nil,
		AuthorityKeyId:              nil,
		OCSPServer:                  nil,
		IssuingCertificateURL:       nil,
		DNSNames:                    nil,
		EmailAddresses:              nil,
		IPAddresses:                 nil,
		URIs:                        nil,
		PermittedEmailAddresses:     nil,
		ExcludedEmailAddresses:      nil,
		CRLDistributionPoints:       nil,
		PolicyIdentifiers:           nil,
	}
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	cert, err := x509.CreateCertificate(rand.Reader, &template, &template, key.Public(), key)
	if err != nil {
		return nil, err
	}
	c, err := x509.ParseCertificate(cert)
	if err != nil {
		return nil, err
	}
	fmt.Println(c.IsCA)
	return &Certificate{
		Certificate: c,
		public:      key.Public(),
		private:     key,
	}, nil
}

// This is NOT a healthy example of an intermediate certificate, this is nothing
// more than a signed certificate with IsCA set to true. Not even any
// basic constraints are defined. Please do not think that this will be
// acceptable to any system, let alone lint particularly well.
func newIntermediate(parent *Certificate) (*Certificate, error) {
	// Edit this template to look like whatever intermediate you need.
	template := x509.Certificate{
		Raw:                         nil,
		RawTBSCertificate:           nil,
		RawSubjectPublicKeyInfo:     nil,
		RawSubject:                  nil,
		RawIssuer:                   nil,
		Signature:                   nil,
		SignatureAlgorithm:          0,
		PublicKeyAlgorithm:          0,
		PublicKey:                   nil,
		Version:                     0,
		SerialNumber:                nextSerial(),
		Issuer:                      pkix.Name{},
		Subject:                     pkix.Name{},
		NotBefore:                   time.Time{},
		NotAfter:                    time.Date(9999, 0, 0, 0, 0, 0, 0, time.UTC),
		KeyUsage:                    0,
		Extensions:                  nil,
		ExtraExtensions:             nil,
		UnhandledCriticalExtensions: nil,
		ExtKeyUsage:                 nil,
		UnknownExtKeyUsage:          nil,
		BasicConstraintsValid:       true,
		IsCA:                        true,
		MaxPathLen:                  0,
		MaxPathLenZero:              false,
		SubjectKeyId:                nil,
		AuthorityKeyId:              nil,
		OCSPServer:                  nil,
		IssuingCertificateURL:       nil,
		DNSNames:                    nil,
		EmailAddresses:              nil,
		IPAddresses:                 nil,
		URIs:                        nil,
		PermittedEmailAddresses:     nil,
		ExcludedEmailAddresses:      nil,
		CRLDistributionPoints:       nil,
		PolicyIdentifiers:           nil,
	}
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	cert, err := x509.CreateCertificate(rand.Reader, &template, parent.Certificate, key.Public(), parent.private)
	if err != nil {
		return nil, err
	}
	c, err := x509.ParseCertificate(cert)
	if err != nil {
		return nil, err
	}
	return &Certificate{
		Certificate: c,
		public:      key.Public(),
		private:     key,
	}, nil
}

// This is NOT a healthy example of a leaf certificate, this is nothing
// more than a self signed certificate with IsCA set to false. Not even any
// basic constraints are defined. Please do not think that this will be
// acceptable to any system, let alone lint particularly well.
func newLeaf(trustAnchor *Certificate, intermediates []*Certificate) (*Certificate, error) {
	var parent *Certificate
	if len(intermediates) == 0 {
		parent = trustAnchor
	} else {
		parent = intermediates[len(intermediates)-1]
	}
	// Edit this template to look like whatever leaf cert you need.
	template := x509.Certificate{
		Raw:                         nil,
		RawTBSCertificate:           nil,
		RawSubjectPublicKeyInfo:     nil,
		RawSubject:                  nil,
		RawIssuer:                   nil,
		Signature:                   nil,
		SignatureAlgorithm:          0,
		PublicKeyAlgorithm:          0,
		PublicKey:                   nil,
		Version:                     0,
		SerialNumber:                nextSerial(),
		Issuer:                      pkix.Name{},
		Subject:                     pkix.Name{},
		NotBefore:                   time.Time{},
		NotAfter:                    time.Date(9999, 0, 0, 0, 0, 0, 0, time.UTC),
		KeyUsage:                    0,
		Extensions:                  nil,
		ExtraExtensions:             nil,
		UnhandledCriticalExtensions: nil,
		ExtKeyUsage:                 nil,
		UnknownExtKeyUsage:          nil,
		BasicConstraintsValid:       false,
		IsCA:                        false,
		MaxPathLen:                  0,
		MaxPathLenZero:              false,
		SubjectKeyId:                nil,
		AuthorityKeyId:              nil,
		OCSPServer:                  nil,
		IssuingCertificateURL:       nil,
		DNSNames:                    nil,
		EmailAddresses:              nil,
		IPAddresses:                 nil,
		URIs:                        nil,
		PermittedEmailAddresses:     nil,
		ExcludedEmailAddresses:      nil,
		CRLDistributionPoints:       nil,
		PolicyIdentifiers:           nil,
	}
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	cert, err := x509.CreateCertificate(rand.Reader, &template, parent.Certificate, key.Public(), parent.private)
	if err != nil {
		return nil, err
	}
	c, err := x509.ParseCertificate(cert)
	if err != nil {
		return nil, err
	}
	return &Certificate{
		Certificate: c,
		public:      key.Public(),
		private:     key,
	}, nil
}

// Formats the given certificate into OpenSSL's textual output. For example:
//
// Certificate:
//    Data:
//        Version: 3 (0x2)
//        Serial Number: 1 (0x1)
//        Signature Algorithm: ecdsa-with-SHA256
//        Issuer:
//        Validity
//            Not Before: Feb 14 17:21:17 2021 GMT
//            Not After : Feb 14 17:21:17 2021 GMT
//        Subject:
//        Subject Public Key Info:
//            Public Key Algorithm: id-ecPublicKey
//                Public-Key: (256 bit)
//                pub:
//                    04:76:2b:19:b8:b4:f4:d9:9e:66:8a:6a:f3:bf:c5:
//                    df:83:43:d6:53:bf:9e:5a:b8:b1:5d:99:8c:4e:d7:
//                    59:25:fd:5c:08:16:23:19:61:c4:cc:c2:f7:db:ac:
//                    72:a5:5e:65:35:f3:64:e2:9b:af:f9:04:c9:99:61:
//                    57:3e:ee:9c:b3
//                ASN1 OID: prime256v1
//                NIST CURVE: P-256
//        X509v3 extensions:
//            X509v3 Subject Key Identifier:
//                6E:3F:50:3A:07:4E:10:AA:74:31:8F:3B:B3:4F:30:96:D3:6F:EF:AE
//    Signature Algorithm: ecdsa-with-SHA256
//         30:44:02:20:11:3f:4a:25:63:10:fa:2d:96:00:e8:23:8c:62:
//         40:c4:8d:31:31:d0:96:f2:7d:28:34:3a:2c:23:9f:bb:28:7e:
//         02:20:1b:8a:68:6d:ef:c4:d7:19:46:48:bf:b0:18:85:31:37:
//         ce:2f:04:27:7c:a3:d2:47:4d:e1:1f:c3:1a:3e:e3:8f
// -----BEGIN CERTIFICATE-----
// MIIBDjCBtqADAgECAgEBMAoGCCqGSM49BAMCMAAwHhcNMjEwMjE0MTcyMTE3WhcN
// MjEwMjE0MTcyMTE3WjAAMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEdisZuLT0
// 2Z5mimrzv8Xfg0PWU7+eWrixXZmMTtdZJf1cCBYjGWHEzML326xypV5lNfNk4puv
// +QTJmWFXPu6cs6MhMB8wHQYDVR0OBBYEFG4/UDoHThCqdDGPO7NPMJbTb++uMAoG
// CCqGSM49BAMCA0cAMEQCIBE/SiVjEPotlgDoI4xiQMSNMTHQlvJ9KDQ6LCOfuyh+
// AiAbimht78TXGUZIv7AYhTE3zi8EJ3yj0kdN4R/DGj7jjw==
// -----END CERTIFICATE-----
//
// Requires a copy of openssl in $PATH as it is simply making a
// subprocess call out to it.
func openSSLFormatCertificate(cert *Certificate) (string, error) {
	block := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	})
	cmd := exec.Command("openssl", "x509", "-text")
	cmd.Stdin = strings.NewReader(string(block))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// nextSerial is a simple, thread safe, sequential serial number generator.
// Serial numbers begin an 1 and monotonically increase with each call.
var nextSerial = func() func() *big.Int {
	l := sync.Mutex{}
	var serial int64
	return func() *big.Int {
		l.Lock()
		defer l.Unlock()
		serial++
		return big.NewInt(serial)
	}
}()

// Uncomment this and use it if you would like to have random serial numbers.
//
//	// nextRandomSerial randomly generates a single serial number. Serial
//	// numbers generated by sequential calls to this function will be related
//	// to each other in any way.
//	func nextRandomSerial() *big.Int {
//		serial, err := rand.Int(rand.Reader, big.NewInt(int64(math.Pow(2, 160))))
//		if err != nil {
//			panic(err)
//		}
//		return serial
//	}

type Certificate struct {
	*x509.Certificate
	public  interface{}
	private interface{}
}
