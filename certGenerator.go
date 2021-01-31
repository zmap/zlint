// based on https://stackoverflow.com/q/26441547/1426535

package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	"time"
)

func main() {

	var filename = "test"

	template := x509.Certificate{
		SerialNumber:   big.NewInt(time.Now().Unix()),
		Subject:        pkix.Name{CommonName: "example"},
		NotBefore:      time.Date(2021, time.January, 31, 0, 0, 0, 0, time.UTC),
		NotAfter:       time.Date(2021, time.January, 31, 0, 0, 0, 0, time.UTC).AddDate(1, 0, 0),
		KeyUsage:       x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:    []x509.ExtKeyUsage{x509.ExtKeyUsageEmailProtection},
		EmailAddresses: []string{"test@example.com"},
	}

	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	crt, err := x509.CreateCertificate(rand.Reader, &template, &template, &privatekey.PublicKey, privatekey)
	if err != nil {
		panic(err)
	}

	var certOut, keyOut bytes.Buffer
	pem.Encode(&certOut, &pem.Block{Type: "CERTIFICATE", Bytes: crt})
	pem.Encode(&keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privatekey)})
	ioutil.WriteFile("c:\\temp\\private_"+filename+".key", keyOut.Bytes(), 0644)
	ioutil.WriteFile("c:\\temp\\certificate_"+filename+".pem", certOut.Bytes(), 0644)
}
