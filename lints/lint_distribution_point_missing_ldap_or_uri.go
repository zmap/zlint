// lint_distribution_point_missing_ldap_or_uri.go
/************************************************
RFC 5280: 4.2.1.13
When present, DistributionPointName SHOULD include at least one LDAP or HTTP URI.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type distribNoLDAPorURI struct{}

func (l *distribNoLDAPorURI) Initialize() error {
	return nil
}

func (l *distribNoLDAPorURI) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.CrlDistOID)
}

func (l *distribNoLDAPorURI) Execute(c *x509.Certificate) *LintResult {
	for _, point := range c.CRLDistributionPoints {
		if point = strings.ToLower(point); strings.HasPrefix(point, "http://") || strings.HasPrefix(point, "ldap://") {
			return &LintResult{Status: Pass}
		}
	}
	return &LintResult{Status: Warn}
}

func init() {
	RegisterLint(&Lint{
		Name:           "w_distribution_point_missing_ldap_or_uri",
		Description:    "When present in the CRLDistributionPoints extension, DistributionPointName SHOULD include at least one LDAP or HTTP URI",
		ReadableSource: "RFC 5280: 4.2.1.13",
		Source:         RFC5280,
		EffectiveDate:  util.RFC5280Date,
		Lint:           &distribNoLDAPorURI{},
	})
}
