package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCAEKUMissing struct {
	// Internal data here
}

func (l *subCAEKUMissing) Initialize() error {
	return nil
}

func (l *subCAEKUMissing) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsSubCA(c)
}

func (l *subCAEKUMissing) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if util.IsExtInCert(c, util.EkuSynOid) {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_ca_eku_missing",
		Description:   "Subordinate CA certificate MUST have extkeyUsage extension",
		Source:        "BRs: 7.1.5",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCAEKUMissing{},
	})
}
