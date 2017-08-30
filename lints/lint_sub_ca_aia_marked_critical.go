package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCaAIAMarkedCritical struct {
	// Internal data here
}

func (l *subCaAIAMarkedCritical) Initialize() error {
	return nil
}

func (l *subCaAIAMarkedCritical) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsSubCA(c) && util.IsExtInCert(c, util.AiaOID)
}

func (l *subCaAIAMarkedCritical) RunTest(c *x509.Certificate) (ResultStruct, error) {
	e := util.GetExtFromCert(c, util.AiaOID)
	if e.Critical {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_ca_aia_marked_critical",
		Description:   "Subordinate CA Certificate: authorityInformationAccess MUST NOT be marked critical",
		Source:        "BRs: 7.1.2.2",
		EffectiveDate: util.ZeroDate,
		Test:          &subCaAIAMarkedCritical{},
	})
}
