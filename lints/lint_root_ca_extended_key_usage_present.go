// lint_root_ca_extended_key_usage_present.go
/************************************************
BRs: 7.1.2.1d extendedKeyUsage
This extension MUST NOT be present.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type rootCAContainsEKU struct {
	// Internal data here
}

func (l *rootCAContainsEKU) Initialize() error {
	return nil
}

func (l *rootCAContainsEKU) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsRootCA(c)
}

func (l *rootCAContainsEKU) RunTest(c *x509.Certificate) (ResultStruct, error) {
	// Add actual lint here
	if util.IsExtInCert(c, util.EkuSynOid) {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_root_ca_extended_key_usage_present",
		Description:   "Root CA Certificate: extendedKeyUsage MUST NOT be present.t",
		Source:        "BRs: 7.1.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &rootCAContainsEKU{},
	})
}
