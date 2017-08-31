// lint_basic_constraints_not_critical.go
/************************************************
RFC 5280: 4.2.1.9
Conforming CAs MUST include this extension in all CA certificates
that contain public keys used to validate digital signatures on
certificates and MUST mark the extension as critical in such
certificates.  This extension MAY appear as a critical or non-
critical extension in CA certificates that contain public keys used
exclusively for purposes other than validating digital signatures on
certificates.  Such CA certificates include ones that contain public
keys used exclusively for validating digital signatures on CRLs and
ones that contain key management public keys used with certificate.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type basicConstCrit struct {
	// Internal data here
}

func (l *basicConstCrit) Initialize() error {
	return nil
}

func (l *basicConstCrit) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return c.IsCA && util.IsExtInCert(c, util.BasicConstOID)
}

func (l *basicConstCrit) RunTest(c *x509.Certificate) (ResultStruct, error) {
	// Add actual lint here
	if e := util.GetExtFromCert(c, util.BasicConstOID); e != nil {
		if e.Critical {
			return ResultStruct{Result: Pass}, nil
		} else {
			return ResultStruct{Result: Error}, nil
		}
	} else {
		return ResultStruct{Result: NA}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_basic_constraints_not_critical",
		Description:   "basicConstraints MUST appear as a critical extension",
		Source:        "RFC 5280: 4.2.1.9",
		EffectiveDate: util.RFC2459Date,
		Test:          &basicConstCrit{},
	})
}
