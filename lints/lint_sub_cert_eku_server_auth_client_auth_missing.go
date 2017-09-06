// lint_sub_cert_eku_server_auth_client_auth_missing.go
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

type subExtKeyUsageClientOrServer struct{}

func (l *subExtKeyUsageClientOrServer) Initialize() error {
	return nil
}

func (l *subExtKeyUsageClientOrServer) CheckApplies(c *x509.Certificate) bool {
	return c.ExtKeyUsage != nil
}

func (l *subExtKeyUsageClientOrServer) Execute(c *x509.Certificate) *LintResult {
	// Add actual lint here
	for _, kp := range c.ExtKeyUsage {
		if kp == x509.ExtKeyUsageServerAuth || kp == x509.ExtKeyUsageClientAuth {
			// If we find either of ServerAuth or ClientAuth, Pass
			return &LintResult{Status: Pass}
		}
	}
	// If neither were found, Error
	return &LintResult{Status: Error}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_eku_server_auth_client_auth_missing",
		Description:   "Subscriber certificates MUST have have either id-kp-serverAuth or id-kp-clientAuth or both present in extKeyUsage",
		Citation:      "BRs: 7.1.2.3",
		Source:        CABFBaselineRequirements,
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &subExtKeyUsageClientOrServer{},
	})
}
