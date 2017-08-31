package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type rootCAKeyUsagePresent struct {
	// Internal data here
}

func (l *rootCAKeyUsagePresent) Initialize() error {
	return nil
}

func (l *rootCAKeyUsagePresent) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsRootCA(c)
}

func (l *rootCAKeyUsagePresent) RunTest(c *x509.Certificate) (ResultStruct, error) {
	// Add actual lint here
	if util.IsExtInCert(c, util.KeyUsageOID) {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_root_ca_key_usage_present",
		Description:   "Root CA certificates MUST have Key Usage Extension Present",
		Source:        "BRs: 7.1.2.1",
		EffectiveDate: util.RFC2459Date,
		Test:          &rootCAKeyUsagePresent{},
	})
}
