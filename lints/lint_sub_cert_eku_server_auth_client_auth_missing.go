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

type subExtKeyUsageClientOrServer struct {
	// Internal data here
}

func (l *subExtKeyUsageClientOrServer) Initialize() error {
	return nil
}

func (l *subExtKeyUsageClientOrServer) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return c.ExtKeyUsage != nil
}

func (l *subExtKeyUsageClientOrServer) RunTest(c *x509.Certificate) (ResultStruct, error) {
	// Add actual lint here
	for _, kp := range c.ExtKeyUsage {
		if kp == x509.ExtKeyUsageServerAuth || kp == x509.ExtKeyUsageClientAuth {
			// If we find either of ServerAuth or ClientAuth, Pass
			return ResultStruct{Result: Pass}, nil
		}
	}
	// If neither were found, Error
	return ResultStruct{Result: Error}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_eku_server_auth_client_auth_missing",
		Description:   "Subscriber certificates MUST have have either id-kp-serverAuth or id-kp-clientAuth or both present in extKeyUsage",
		Source:        "BRs: 7.1.2.3",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subExtKeyUsageClientOrServer{},
	})
}
