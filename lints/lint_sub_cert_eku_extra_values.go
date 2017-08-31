// lint_sub_cert_eku_extra_values.go
/*******************************************************************************************************
BRs: 7.1.2.3
extKeyUsage (required)
Either the value id-kp-serverAuth [RFC5280] or id-kp-clientAuth [RFC5280] or both values MUST be present. id-kp-emailProtection [RFC5280] MAY be present. Other values SHOULD NOT be present.
*******************************************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subExtKeyUsageLegalUsage struct {
	// Internal data here
}

func (l *subExtKeyUsageLegalUsage) Initialize() error {
	return nil
}

func (l *subExtKeyUsageLegalUsage) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return c.ExtKeyUsage != nil
}

func (l *subExtKeyUsageLegalUsage) RunTest(c *x509.Certificate) (ResultStruct, error) {
	// Add actual lint here
	for _, kp := range c.ExtKeyUsage {
		if kp == x509.ExtKeyUsageServerAuth ||
			kp == x509.ExtKeyUsageClientAuth ||
			kp == x509.ExtKeyUsageEmailProtection {
			// If we find any of these three, considered passing, continue
			continue
		} else {
			// A bad usage was found, report and leave
			return ResultStruct{Result: Warn}, nil
		}
	}
	// If no bad usage was found, pass
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_sub_cert_eku_extra_values",
		Description:   "Subscriber Certificate: extKeyUsage either the value id-kp-serverAuth or id-kp-clientAuth or both values MUST be present.",
		Source:        "BRs: 7.1.2.3",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subExtKeyUsageLegalUsage{},
	})
}
