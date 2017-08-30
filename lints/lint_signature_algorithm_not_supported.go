package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type signatureAlgorithmNotSupported struct {
	// Internal data here
}

func (l *signatureAlgorithmNotSupported) Initialize() error {
	return nil
}

func (l *signatureAlgorithmNotSupported) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *signatureAlgorithmNotSupported) RunTest(c *x509.Certificate) (ResultStruct, error) {

	if c.SignatureAlgorithm == x509.SHA1WithRSA || c.SignatureAlgorithm == x509.SHA256WithRSA || c.SignatureAlgorithm == x509.SHA384WithRSA || c.SignatureAlgorithm == x509.SHA512WithRSA || c.SignatureAlgorithm == x509.DSAWithSHA1 || c.SignatureAlgorithm == x509.DSAWithSHA256 || c.SignatureAlgorithm == x509.ECDSAWithSHA1 || c.SignatureAlgorithm == x509.ECDSAWithSHA256 || c.SignatureAlgorithm == x509.ECDSAWithSHA384 || c.SignatureAlgorithm == x509.ECDSAWithSHA512 {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_signature_algorithm_not_supported",
		Description:   "Certificates MUST meet the following requirements for algorithm type: SHA-1*, SHA-256, SHA-384, SHA-512",
		Source:        "BRs: 6.1.5",
		EffectiveDate: util.ZeroDate,
		Test:          &signatureAlgorithmNotSupported{},
	})
}
