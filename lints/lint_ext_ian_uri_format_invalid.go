/************************************************
The name MUST include both a
scheme (e.g., "http" or "ftp") and a scheme-specific-part.
************************************************/

package lints

import (
	"net/url"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type IANURIFormat struct{}

func (l *IANURIFormat) Initialize() error {
	return nil
}

func (l *IANURIFormat) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.IssuerAlternateNameOID)
}

func (l *IANURIFormat) Execute(c *x509.Certificate) *LintResult {
	for _, uri := range c.IANURIs {
		parsed_uri, err := url.Parse(uri)

		if err != nil {
			return &LintResult{Status: Error}
		}

		//scheme
		if parsed_uri.Scheme == "" {
			return &LintResult{Status: Error}
		}

		//scheme-specific part
		if parsed_uri.Host == "" && parsed_uri.User == nil && parsed_uri.Opaque == "" && parsed_uri.Path == "" {
			return &LintResult{Status: Error}
		}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_ian_uri_format_invalid",
		Description:   "URIs in the subjectAltName extension MUST have a scheme and scheme specific part",
		Citation:      "RFC5280: 4.2.1.6",
		Source:        RFC5280,
		EffectiveDate: util.RFC5280Date,
		Lint:          &IANURIFormat{},
	})
}
