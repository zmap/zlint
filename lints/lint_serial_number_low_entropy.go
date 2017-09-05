package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type serialNumberLowEntropy struct{}

func (l *serialNumberLowEntropy) Initialize() error {
	return nil
}

func (l *serialNumberLowEntropy) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *serialNumberLowEntropy) Execute(c *x509.Certificate) *LintResult {
	if len(c.SerialNumber.Bytes()) < 8 {
		return &LintResult{Status: Warn}
	} else {
		return &LintResult{Status: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_serial_number_low_entropy",
		Description:   "Effective September 30, 2016, CAs SHALL generate non‐sequential Certificate serial numbers greater than zero (0) containing at least 64 bits of output from a CSPRNG.",
		Source:        "BRs: 7.1",
		Type:          BRs,
		EffectiveDate: util.CABSerialNumberEntropyDate,
		Lint:          &serialNumberLowEntropy{},
	})
}
