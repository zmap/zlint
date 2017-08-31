// lint_inhibit_any_policy_not_critical.go
/************************************************
4.2.1.14.  Inhibit anyPolicy
   The inhibit anyPolicy extension can be used in certificates issued to CAs.
   The inhibit anyPolicy extension indicates that the special anyPolicy OID,
   with the value { 2 5 29 32 0 }, is not considered an explicit match for other
   certificate policies except when it appears in an intermediate self-issued
   CA certificate. The value indicates the number of additional non-self-issued
   certificates that may appear in the path before anyPolicy is no longer permitted.
   For example, a value of one indicates that anyPolicy may be processed in
   certificates issued by the subject of this certificate, but not in additional
   certificates in the path.

   Conforming CAs MUST mark this extension as critical.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type InhibitAnyPolicyNotCritical struct {
	// Internal data here
}

func (l *InhibitAnyPolicyNotCritical) Initialize() error {
	return nil
}

func (l *InhibitAnyPolicyNotCritical) CheckApplies(cert *x509.Certificate) bool {
	return util.IsExtInCert(cert, util.InhibitAnyPolicyOID)
}

func (l *InhibitAnyPolicyNotCritical) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	if anyPol := util.GetExtFromCert(cert, util.InhibitAnyPolicyOID); !anyPol.Critical {
		return ResultStruct{Result: Error}, nil
	} //else
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_inhibit_any_policy_not_critical",
		Description:   "CAs MUST mark the inhibitAnyPolicy extension as critical",
		Source:        "RFC 5280: 4.2.1.14",
		EffectiveDate: util.RFC3280Date,
		Test:          &InhibitAnyPolicyNotCritical{},
	})
}
