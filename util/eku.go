package util

import (
	"github.com/zmap/zcrypto/x509"
)

func IsTestableBRCertificate(cert *x509.Certificate) bool {
	if len(cert.ExtKeyUsage) == 0 {
		return true
	}
	for _, eku := range cert.ExtKeyUsage {
		if eku == x509.ExtKeyUsageAny || eku == x509.ExtKeyUsageServerAuth {
			return true
		}
	}
	return false
}
