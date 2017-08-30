// lint_ext_ian_critical.go
/************************************************
Issuer Alternative Name
   As with Section 4.2.1.6, this extension is used to associate Internet style identities with the certificate issuer. Issuer alternative name MUST be encoded as in 4.2.1.6.  Issuer alternative names are not processed as part of the certification path validation algorithm in Section 6. (That is, issuer alternative names are not used in name chaining and name constraints are not enforced.)
   Where present, conforming CAs SHOULD mark this extension as non-critical.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type ExtIANCritical struct {
	// Internal data here
}

func (l *ExtIANCritical) Initialize() error {
	return nil
}

func (l *ExtIANCritical) CheckApplies(cert *x509.Certificate) bool {
	return util.IsExtInCert(cert, util.IssuerAlternateNameOID)
}

func (l *ExtIANCritical) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	if util.GetExtFromCert(cert, util.IssuerAlternateNameOID).Critical {
		return ResultStruct{Result: Warn}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_ext_ian_critical",
		Description:   "Issuer alternate name should be marked as non-critical",
		Source:        "RFC 5280: 4.2.1.7",
		EffectiveDate: util.RFC2459Date,
		Test:          &ExtIANCritical{},
	})
}
