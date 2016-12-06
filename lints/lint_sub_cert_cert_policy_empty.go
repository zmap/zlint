// lint_sub_cert_cert_policy_empty.go
/********************************************************************************************************************
CAB: 7.1.6.4
Subscriber Certificates
A Certificate issued to a Subscriber MUST contain one or more policy identifier(s), defined by the Issuing CA, in
the Certificate’s certificatePolicies extension that indicates adherence to and compliance with these Requirements.
CAs complying with these Requirements MAY also assert one of the reserved policy OIDs in such Certificates.
*********************************************************************************************************************/

package lints

import (

	"github.com/teamnsrg/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
)

type subCertPolicyEmpty struct {
	// Internal data here
}

func (l *subCertPolicyEmpty) Initialize() error {
	return nil
}

func (l *subCertPolicyEmpty) CheckApplies(c *x509.Certificate) bool {
	return !util.IsCaCert(c)
}

func (l *subCertPolicyEmpty) RunTest(c *x509.Certificate) (ResultStruct, error) {
	// Add actual lint here
	if util.IsExtInCert(c, util.CertPolicyOID) && c.PolicyIdentifiers != nil {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Warn}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "sub_cert_cert_policy_empty",
		Description:   "Subscriber certificates must contain at least one policy identifier that indicates adherance to CAB standards",
		Providence:    "CAB: 7.1.6.4",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCertPolicyEmpty{}})
}
