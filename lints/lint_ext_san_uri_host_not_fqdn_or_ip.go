// lint_ext_san_uri_host_not_fqdn_or_ip.go
/*********************************************************************
When the subjectAltName extension contains a URI, the name MUST be
stored in the uniformResourceIdentifier (an IA5String).  The name
MUST NOT be a relative URI, and it MUST follow the URI syntax and
encoding rules specified in [RFC3986].  The name MUST include both a
scheme (e.g., "http" or "ftp") and a scheme-specific-part.  URIs that
include an authority ([RFC3986], Section 3.2) MUST include a fully
qualified domain name or IP address as the host.  Rules for encoding
Internationalized Resource Identifiers (IRIs) are specified in
Section 7.4.
*********************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SANURIHost struct{}

func (l *SANURIHost) Initialize() error {
	return nil
}

func (l *SANURIHost) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *SANURIHost) Execute(c *x509.Certificate) *LintResult {
	for _, uri := range c.URIs {
		if uri != "" {
			host := util.GetHost(uri)
			if !util.AuthIsFQDNOrIP(host) {
				return &LintResult{Status: Error}
			}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:           "e_ext_san_uri_host_not_fqdn_or_ip",
		Description:    "URIs that include an authority ([RFC3986], Section 3.2) MUST include a fully qualified domain name or IP address as the host",
		ReadableSource: "RFC 5280: 4.2.1.7",
		Source:         RFC5280,
		EffectiveDate:  util.RFC5280Date,
		Lint:           &SANURIHost{},
	})
}
