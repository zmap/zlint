package lints

import (
	"encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertNotCA struct{}

func (l *subCertNotCA) Initialize() error {
	return nil
}

func (l *subCertNotCA) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.KeyUsageOID) && c.KeyUsage&x509.KeyUsageCertSign == 0 && util.IsExtInCert(c, util.BasicConstOID)
}

func (l *subCertNotCA) Execute(c *x509.Certificate) *LintResult {
	e := util.GetExtFromCert(c, util.BasicConstOID)
	var constraints basicConstraints
	if _, err := asn1.Unmarshal(e.Value, &constraints); err != nil {
		return &LintResult{Status: Fatal}
	}
	if constraints.IsCA == true {
		return &LintResult{Status: Error}
	} else {
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:           "e_sub_cert_not_is_ca",
		Description:    "Subscriber Certificate: basicContrainsts cA field MUST NOT be true.",
		ReadableSource: "BRs: 7.1.2.3",
		Source:         CABFBaselineRequirements,
		EffectiveDate:  util.CABEffectiveDate,
		Lint:           &subCertNotCA{},
	})
}
