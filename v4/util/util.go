package util

import "github.com/zmap/zcrypto/x509"

func IsCA(c *x509.Certificate) bool {
	return c.BasicConstraintsValid && c.IsCA
}

func HasTBSCertificate(c *x509.Certificate) bool {
	return c != nil
}
