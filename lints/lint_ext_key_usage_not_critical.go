// lint_ext_key_usage_not_critical.go
// "When present, conforming CAs SHOULD mark this extension as critical."

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type checkKeyUsageCritical struct{}

func (l *checkKeyUsageCritical) Initialize() error {
	return nil
}

func (l *checkKeyUsageCritical) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.KeyUsageOID)
}

func (l *checkKeyUsageCritical) Execute(c *x509.Certificate) ResultStruct {
	// Add actual lint here
	keyUsage := util.GetExtFromCert(c, util.KeyUsageOID)
	if keyUsage == nil {
		return ResultStruct{Result: NA}
	}
	if keyUsage.Critical {
		return ResultStruct{Result: Pass}
	} else {
		return ResultStruct{Result: Warn}
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_ext_key_usage_not_critical",
		Description:   "The keyUsage extension SHOULD be critical",
		Source:        "RFC 5280: 4.2.1.3",
		EffectiveDate: util.RFC2459Date,
		Lint:          &checkKeyUsageCritical{},
	})
}
