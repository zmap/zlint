// lint_sub_cert_cert_policy_empty.go
/********************************************************************************************************************
BRs: 7.1.6.4
Subscriber Certificates
A Certificate issued to a Subscriber MUST contain one or more policy identifier(s), defined by the Issuing CA, in
the Certificate’s certificatePolicies extension that indicates adherence to and complIANce with these Requirements.
CAs complying with these Requirements MAY also assert one of the reserved policy OIDs in such Certificates.
*********************************************************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertPolicyEmpty struct {
	// Internal data here
}

func (l *subCertPolicyEmpty) Initialize() error {
	return nil
}

func (l *subCertPolicyEmpty) CheckApplies(c *x509.Certificate) bool {
	return !util.IsCACert(c)
}

func (l *subCertPolicyEmpty) RunTest(c *x509.Certificate) (ResultStruct, error) {
	// Add actual lint here
	if util.IsExtInCert(c, util.CertPolicyOID) && c.PolicyIdentifiers != nil {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_cert_policy_empty",
		Description:   "Subscriber certificates must contain at least one policy identifier that indicates adherence to CAB standards",
		Source:        "BRs: 7.1.6.4",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCertPolicyEmpty{},
	})
}
