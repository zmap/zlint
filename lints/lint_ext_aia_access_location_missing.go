// lint_ext_aia_access_location_missing.go
/************************************************
RFC 5280: 4.2.2.1
An authorityInfoAccess extension may include multiple instances of
   the id-ad-caIssuers accessMethod.  The different instances may
   specify different methods for accessing the same information or may
   point to different information.  When the id-ad-caIssuers
   accessMethod is used, at least one instance SHOULD specify an
   accessLocation that is an HTTP [RFC2616] or LDAP [RFC4516] URI.

************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type aiaNoHTTPorLDAP struct {
	// Internal data here
}

func (l *aiaNoHTTPorLDAP) Initialize() error {
	return nil
}

func (l *aiaNoHTTPorLDAP) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return util.IsExtInCert(c, util.AiaOID) && c.IssuingCertificateURL != nil
}

func (l *aiaNoHTTPorLDAP) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, caIssuer := range c.IssuingCertificateURL {
		if caIssuer = strings.ToLower(caIssuer); strings.HasPrefix(caIssuer, "http://") || strings.HasPrefix(caIssuer, "ldap://") {
			return ResultStruct{Result: Pass}, nil
		}
	}
	return ResultStruct{Result: Warn}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "w_ext_aia_access_location_missing",
		Description:   "When the id-ad-caIssuers accessMethod is used, at least one instance SHOULD specify an accessLocation that is an HTTP or LDAP URI",
		Source:        "RFC 5280: 4.2.2.1",
		EffectiveDate: util.RFC5280Date,
		Test:          &aiaNoHTTPorLDAP{},
	})
}
