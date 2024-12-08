package main

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

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/pem"
	"fmt"
	"math/big"
	"os/exec"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/zmap/zcrypto/encoding/asn1"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/v3/util"
)

type certificateList struct {
	TBSCertList        tbsCertificateList
	SignatureAlgorithm pkix.AlgorithmIdentifier
	SignatureValue     asn1.BitString
}
type tbsCertificateList struct {
	Raw                 asn1.RawContent
	Version             int `asn1:"optional,default:0"`
	Signature           pkix.AlgorithmIdentifier
	Issuer              asn1.RawValue
	ThisUpdate          time.Time
	NextUpdate          time.Time                 `asn1:"optional"`
	RevokedCertificates []pkix.RevokedCertificate `asn1:"optional"`
	Extensions          []pkix.Extension          `asn1:"tag:0,optional,explicit"`
}

func main() {
	//	CN=Root CA, C=US, O=Example Root CA
	issuer := pkix.Name{
		CommonName:   "Root CA",
		Country:      []string{"US"},
		Organization: []string{"Example Root CA"},
	}
	tbs := tbsCertificateList{
		Version: 1,
		Signature: pkix.AlgorithmIdentifier{
			Algorithm:  util.OidSHA256WithRSAEncryption,
			Parameters: asn1.NullRawValue,
		},
		Issuer:     asn1.RawValue{Tag: asn1.TagSequence, FullBytes: encode(issuer.ToRDNSequence())},
		ThisUpdate: time.Now(),
		NextUpdate: time.Now().Add(365 * 24 * time.Hour),
		RevokedCertificates: []pkix.RevokedCertificate{
			{
				SerialNumber:   big.NewInt(1), // Revoked Certificate Serial Number
				RevocationTime: time.Now(),
				Extensions: []pkix.Extension{
					{
						Id:    util.ReasonCodeOID,
						Value: encode(asn1.RawValue{Tag: asn1.TagEnum, Bytes: []byte{0x05}}), // reason code "cessationOfOperation" (0x05)
					},
				},
			},
		},
		Extensions: []pkix.Extension{
			{
				Id: util.AuthkeyOID,
				Value: encode(asn1.RawValue{
					Tag:        asn1.TagSequence,
					IsCompound: true,
					Bytes:      []byte{0x80, 0x14, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55}, // Dummy Authority Key Identifier
				}),
			},
		},
	}

	// Marshal the tbsCertList to get the "to-be-signed" bytes
	tbsBytes := encode(tbs)
	signature, err := generateSignature(tbsBytes)
	if err != nil {
		log.Fatalf("Failed to generate signature: %v", err)
	}

	// Create the final CRL
	crl := certificateList{
		TBSCertList: tbs,
		SignatureAlgorithm: pkix.AlgorithmIdentifier{
			Algorithm:  util.OidSHA256WithRSAEncryption,
			Parameters: asn1.NullRawValue,
		},
		SignatureValue: asn1.BitString{
			Bytes:     signature,
			BitLength: len(signature) * 8,
		},
	}

	// Marshal the complete CRL
	crlBytes, err := asn1.Marshal(crl)
	if err != nil {
		log.Fatalf("Failed to marshal CRL: %v", err)
	}
	pem := pem.EncodeToMemory(&pem.Block{Type: "X509 CRL", Bytes: crlBytes})
	fmted, err := openSSLFormatCRL(pem)
	if err != nil {
		log.Fatalf("Failed to format CRL: %v", err)
	}
	fmt.Println(fmted)
}

func encode(val any) []byte {
	b, err := asn1.Marshal(val)
	if err != nil {
		panic(err)
	}
	return b
}

// Generates a signature for the given tbsBytes using a newly generated RSA key.
func generateSignature(tbsBytes []byte) ([]byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("failed to generate RSA key: %w", err)
	}
	hashed := sha256.Sum256(tbsBytes)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, fmt.Errorf("failed to sign hash: %w", err)
	}
	return signature, nil
}

// Formats the given CRL into OpenSSL's textual output. For example:
// Certificate Revocation List (CRL):
//         Version 2 (0x1)
//         Signature Algorithm: sha256WithRSAEncryption
//         Issuer: CN = Root CA, O = Example Root CA, C = US
//         Last Update: Oct 29 19:00:08 2024
//         Next Update: Oct 29 19:00:08 2025
//         CRL extensions:
//             X509v3 CRL Number:
//                 1
//             X509v3 Authority Key Identifier:
//                 keyid:55:55:55:55:55:55:55:55:55:55:55:55:55:55:55:55:55:55:55:55

// Revoked Certificates:
//
//	Serial Number: 01
//	    Revocation Date: Oct 29 19:00:08 2024
//	    CRL entry extensions:
//	        X509v3 CRL Reason Code:
//	            Cessation Of Operation
//	Signature Algorithm: sha256WithRSAEncryption
//	     7a:99:ee:ac:29:7c:6d:01:1f:ba:2f:ef:1a:6b:bc:cf:b1:78:
//	     7b:d4:be:22:fe:3b:a8:20:e7:c7:d8:ad:80:74:ab:75:07:78:
//	     1d:22:28:05:8e:f9:a8:a9:4e:6d:1c:e4:ad:fa:76:cc:a6:46:
//	     e0:fb:39:33:d2:4f:f2:77:8b:8d:b0:80:92:9d:4a:b8:f6:da:
//	     6c:69:dc:b6:1c:1a:55:d1:e3:ed:30:46:88:05:3e:82:92:b9:
//	     00:bc:8c:bb:f3:b6:cb:d3:81:ed:8a:e3:27:12:9e:8c:51:7d:
//	     05:ec:41:a4:a4:28:b7:5f:fb:a1:ce:d2:7f:18:92:74:1f:de:
//	     95:33:a1:35:40:6f:2c:c5:74:20:ea:76:bc:7b:8a:ca:02:02:
//	     8b:18:3a:14:5a:21:f3:f9:15:65:89:46:b3:47:da:17:e6:84:
//	     f7:ad:f3:42:49:8f:f0:ff:a1:2f:b7:20:33:fa:29:47:42:ea:
//	     05:be:30:ce:83:24:2d:ca:cd:07:8b:69:97:5b:36:42:18:2e:
//	     16:93:b3:0c:40:f1:99:59:60:01:49:35:27:8a:3e:2c:f4:88:
//	     ed:95:58:db:bb:2e:e9:de:fe:2f:60:3f:15:ef:40:59:e6:9e:
//	     35:e1:ee:46:21:cd:0a:e4:91:00:e1:be:00:15:60:1f:ca:94:
//	     1f:40:f3:5a
//
// -----BEGIN X509 CRL-----
// MIIB4zCBzAIBATANBgkqhkiG9w0BAQsFADA5MRAwDgYDVQQDEwdSb290IENBMRgw
// FgYDVQQKEw9FeGFtcGxlIFJvb3QgQ0ExCzAJBgNVBAYTAlVTFxEyNDEwMjkyMDAw
// MDgrMDEwMBcRMjUxMDI5MjAwMDA4KzAxMDAwJjAkAgEBFxEyNDEwMjkyMDAwMDgr
// MDEwMDAMMAoGA1UdFQQDCgEFoC8wLTAKBgNVHRQEAwIBATAfBgNVHSMEGDAWgBRV
// VVVVVVVVVVVVVVVVVVVVVVVVVTANBgkqhkiG9w0BAQsFAAOCAQEAepnurCl8bQEf
// ui/vGmu8z7F4e9S+Iv47qCDnx9itgHSrdQd4HSIoBY75qKlObRzkrfp2zKZG4Ps5
// M9JP8neLjbCAkp1KuPbabGncthwaVdHj7TBGiAU+gpK5ALyMu/O2y9OB7YrjJxKe
// jFF9BexBpKQot1/7oc7SfxiSdB/elTOhNUBvLMV0IOp2vHuKygICixg6FFoh8/kV
// ZYlGs0faF+aE963zQkmP8P+hL7cgM/opR0LqBb4wzoMkLcrNB4tpl1s2QhguFpOz
// DEDxmVlgAUk1J4o+LPSI7ZVY27su6d7+L2A/Fe9AWeaeNeHuRiHNCuSRAOG+ABVg
// H8qUH0DzWg==
// -----END X509 CRL-----
//
// Requires a copy of openssl in $PATH as it is simply making a
// subprocess call out to it.
func openSSLFormatCRL(pem []byte) (string, error) {
	cmd := exec.Command("openssl", "crl", "-text")
	cmd.Stdin = strings.NewReader(string(pem))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
