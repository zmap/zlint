package util

import "github.com/zmap/zcrypto/x509"

func IsMailboxValidatedCertificate(c *x509.Certificate) bool {
	for _, oid := range c.PolicyIdentifiers {
		if oid.Equal(SMIMEBRMailboxValidatedLegacyOID) || oid.Equal(SMIMEBRMailboxValidatedMultipurposeOID) || oid.Equal(SMIMEBRMailboxValidatedStrictOID) {
			return true
		}
	}

	return false
}
