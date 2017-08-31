package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCaAIAMarkedCritical struct{}

func (l *subCaAIAMarkedCritical) Initialize() error {
	return nil
}

func (l *subCaAIAMarkedCritical) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubCA(c) && util.IsExtInCert(c, util.AiaOID)
}

func (l *subCaAIAMarkedCritical) Execute(c *x509.Certificate) ResultStruct {
	e := util.GetExtFromCert(c, util.AiaOID)
	if e.Critical {
		return ResultStruct{Result: Error}
	} else {
		return ResultStruct{Result: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_ca_aia_marked_critical",
		Description:   "Subordinate CA Certificate: authorityInformationAccess MUST NOT be marked critical",
		Source:        "BRs: 7.1.2.2",
		EffectiveDate: util.ZeroDate,
		Lint:          &subCaAIAMarkedCritical{},
	})
}
