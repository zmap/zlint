package lints

import (
	"encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type caIsCA struct{}

type basicConstraints struct {
	IsCA       bool `asn1:"optional"`
	MaxPathLen int  `asn1:"optional,default:-1"`
}

func (l *caIsCA) Initialize() error {
	return nil
}

func (l *caIsCA) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.KeyUsageOID) && c.KeyUsage&x509.KeyUsageCertSign != 0 && util.IsExtInCert(c, util.BasicConstOID)
}

func (l *caIsCA) Execute(c *x509.Certificate) LintResult {
	e := util.GetExtFromCert(c, util.BasicConstOID)
	var constraints basicConstraints
	_, err := asn1.Unmarshal(e.Value, &constraints)
	if err != nil {
		return &LintResult{Status: Fatal}
	}
	if constraints.IsCA == true {
		return &LintResult{Status: Pass}
	} else {
		return &LintResult{Status: Error}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ca_is_ca",
		Description:   "Root and Sub CA Certificate: The CA field MUST be set to true.",
		Source:        "BRs: 7.1.2.1, BRs: 7.1.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &caIsCA{},
	})
}
