// lint_root_ca_extended_key_usage_present.go
/************************************************
CAB: 7.1.2.1d extendedKeyUsage
This extension MUST NOT be present.
************************************************/

package lints

import (

	"github.com/teamnsrg/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
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
		Name:          "root_ca_extended_key_usage_present",
		Description:   "Root CA certificates must not have the extendedKeyUsage extension present",
		Providence:    "CAB: 7.1.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &rootCAContainsEKU{}})
}
