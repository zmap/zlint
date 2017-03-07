// lint_ev_serial_number_missing.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type evSNMissing struct {
	// Internal data here
}

func (l *evSNMissing) Initialize() error {
	return nil
}

func (l *evSNMissing) CheckApplies(c *x509.Certificate) bool {
	return util.IsEv(c.PolicyIdentifiers)
}

func (l *evSNMissing) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.SerialNumber.BitLen() == 0 {
		return ResultStruct{Result: Error}, nil
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "ev_serial_number_missing",
		Description:   "EV certificates must include serialNumber in subject",
		Providence:    "",
		EffectiveDate: util.ZeroDate,
		Test:          &evSNMissing{}})
}
