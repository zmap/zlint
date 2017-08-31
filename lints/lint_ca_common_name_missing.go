package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type caCommonNameMissing struct{}

func (l *caCommonNameMissing) Initialize() error {
	return nil
}

func (l *caCommonNameMissing) CheckApplies(c *x509.Certificate) bool {
	return util.IsCACert(c)
}

func (l *caCommonNameMissing) Execute(c *x509.Certificate) ResultStruct {
	if c.Subject.CommonName == "" {
		return ResultStruct{Result: Error}
	} else {
		return ResultStruct{Result: Pass}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ca_common_name_missing",
		Description:   "CA Certificates common name MUST be included.",
		Source:        "BRs: 7.1.4.3.1",
		EffectiveDate: util.RFC2459Date,
		Lint:          &caCommonNameMissing{},
	})
}
