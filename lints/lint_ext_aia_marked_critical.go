// lint_ext_aia_marked_critical.go
/************************************************
Authority Information Access
   The authority information access extension indicates how to access information and services for the issuer of the certificate in which the extension appears. Information and services may include on-line validation services and CA policy data. (The location of CRLs is not specified in this extension; that information is provided by the cRLDistributionPoints extension.) This extension may be included in end entity or CA certificates. Conforming CAs MUST mark this extension as non-critical.
************************************************/
//See also: BRs: 7.1.2.3 & CAB: 7.1.2.2

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type ExtAiaMarkedCritical struct {
	// Internal data here
}

func (l *ExtAiaMarkedCritical) Initialize() error {
	return nil
}

func (l *ExtAiaMarkedCritical) CheckApplies(cert *x509.Certificate) bool {
	return util.IsExtInCert(cert, util.AiaOID)
}

func (l *ExtAiaMarkedCritical) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	if util.GetExtFromCert(cert, util.AiaOID).Critical {
		return ResultStruct{Result: Error}, nil
	} else {
		return ResultStruct{Result: Pass}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_aia_marked_critical",
		Description:   "Conforming CAs must mark the Authority Information Access extension as non-critical",
		Source:        "RFC 5280: 4.2.2.1",
		EffectiveDate: util.RFC2459Date,
		Test:          &ExtAiaMarkedCritical{},
	})
}
