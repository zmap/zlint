// lint_ev_valid_time_too_long.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertValidTimeTooLong struct {
	// Internal data here
}

func (l *subCertValidTimeTooLong) Initialize() error {
	return nil
}

func (l *subCertValidTimeTooLong) CheckApplies(c *x509.Certificate) bool {
	return util.IsSubscriberCert(c)
}

func (l *subCertValidTimeTooLong) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.NotBefore.AddDate(0, 39, 0).Before(c.NotAfter) {
		return ResultStruct{Result: Error}, nil
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_valid_time_too_long",
		Description:   "CAs MUST NOT issue subscriber certificates with validity periods longer than 39 months regardless of circumstance.",
		Source:        "BRs: 6.3.2",
		EffectiveDate: util.SubCert39Month,
		Test:          &subCertValidTimeTooLong{},
	})
}
