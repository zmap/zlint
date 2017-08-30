// lint_ext_ian_uri_host_not_fqdn_or_ip.go
/*********************************************************************
When the issuerAltName extension contains a URI, the name MUST be
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
	"net/url"
)

type IANURIFQDNOrIP struct {
	// Internal data here
}

func (l *IANURIFQDNOrIP) Initialize() error {
	return nil
}

func (l *IANURIFQDNOrIP) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerAlternateNameOID)
}

func (l *IANURIFQDNOrIP) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, uri := range c.IANURIs {
		if uri != "" {
			parsedUrl, err := url.Parse(uri)
			if err != nil {
				return ResultStruct{Result: Error}, nil
			}
			host := parsedUrl.Host
			if !util.AuthIsFQDNOrIP(host) {
				return ResultStruct{Result: Error}, nil
			}
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_ian_uri_host_not_fqdn_or_ip",
		Description:   "URIs that include an authority ([RFC3986], Section 3.2) MUST include a fully qualified domain name or IP address as the host",
		Source:        "RFC 5280: 4.2.1.6",
		EffectiveDate: util.RFC5280Date,
		Test:          &IANURIFQDNOrIP{},
	})
}
