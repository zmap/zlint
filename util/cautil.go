// cautil.go
// contains helper functions to determine if something is a ca or root ca

package util

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
)

func IsCaCert(cert *x509.Certificate) bool {
	return cert.IsCA
}

func IsRootCA(cert *x509.Certificate) bool {
	return IsCaCert(cert) && IsSelfSigned(cert)
}

func IsSubCA(cert *x509.Certificate) bool {
	return IsCaCert(cert) && !IsSelfSigned(cert)
}

func IsSelfSigned(cert *x509.Certificate) bool {
	err := cert.CheckSignature(cert.SignatureAlgorithm, cert.RawTBSCertificate, cert.Signature)
	return err == nil
}

func NotAllNameFieldsAreEmpty(name *pkix.Name) bool {
	//Return true if at least one field is non-empty
	return len(name.Names) >= 1
}
