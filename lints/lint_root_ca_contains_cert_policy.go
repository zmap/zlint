// lint_root_ca_contains_cert_policy.go
/************************************************
CAB: 7.1.2.1c certificatePolicies
This extension SHOULD NOT be present.
************************************************/

package lints

import (

	"github.com/teamnsrg/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
)

type rootCAContainsCertPolicy struct {
	// Internal data here
}

func (l *rootCAContainsCertPolicy) Initialize() error {
	return nil
}

func (l *rootCAContainsCertPolicy) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsRootCA(c)
}

func (l *rootCAContainsCertPolicy) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if util.IsExtInCert(c, util.CertPolicyOID) {
		return ResultStruct{Result: Warn}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "root_ca_contains_cert_policy",
		Description:   "Root CA certs SHOULD NOT contain the certificat policies extension.",
		Providence:    "CAB: 7.1.2.1",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &rootCAContainsCertPolicy{}})
}
