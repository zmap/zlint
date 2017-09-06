// lint_ev_serial_number_missing.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type evSNMissing struct{}

func (l *evSNMissing) Initialize() error {
	return nil
}

func (l *evSNMissing) CheckApplies(c *x509.Certificate) bool {
	return util.IsEV(c.PolicyIdentifiers)
}

func (l *evSNMissing) Execute(c *x509.Certificate) *LintResult {
	if c.SerialNumber.BitLen() == 0 {
		return &LintResult{Status: Error}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ev_serial_number_missing",
		Description:   "EV certificates must include serialNumber in subject",
		Citation:      "BRs: 7.1.6.1",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.ZeroDate,
		Lint:          &evSNMissing{},
	})
}
