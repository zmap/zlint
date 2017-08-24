package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type serialNumberLowEntropy struct {
	// Internal data here
}

func (l *serialNumberLowEntropy) Initialize() error {
	return nil
}

func (l *serialNumberLowEntropy) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return true
}

func (l *serialNumberLowEntropy) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if len(c.SerialNumber.Bytes()) < 8 {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_serial_number_low_entropy",
		Description:   "Effective September 30, 2016, CAs SHALL generate non‐sequential Certificate serial numbers greater than zero (0) containing at least 64 bits of output from a CSPRNG.",
		Provenance:    "BRs: 7.1",
		EffectiveDate: util.CABSerialNumberEntropyDate,
		Test:          &serialNumberLowEntropy{},
		updateReport:  func(report *LintReport, result ResultStruct) { report.ESerialNumberLowEntropy = result },
	})
}
