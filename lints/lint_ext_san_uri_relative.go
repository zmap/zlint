// lint_ext_san_uri_relative.go
/*************************************************************************
When the subjectAltName extension contains a URI, the name MUST be
stored in the uniformResourceIdentifier (an IA5String).  The name
MUST NOT be a relative URI, and it MUST follow the URI syntax and
encoding rules specified in [RFC3986].  The name MUST include both a
scheme (e.g., "http" or "ftp") and a scheme-specific-part.  URIs that
include an authority ([RFC3986], Section 3.2) MUST include a fully
qualified domain name or IP address as the host.  Rules for encoding
Internationalized Resource Identifiers (IRIs) are specified in
Section 7.4.
*************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"net/url"
)

type extSANURIRelative struct {
	// Internal data here
}

func (l *extSANURIRelative) Initialize() error {
	return nil
}

func (l *extSANURIRelative) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *extSANURIRelative) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, uri := range c.URIs {
		parsed_uri, err := url.Parse(uri)

		if err != nil {
			return ResultStruct{Result: Error}, nil
		}

		if !parsed_uri.IsAbs() {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_uri_relative",
		Description:   "When the subjectAlternateName extension is present and a URI is used, the name MUST NOT be a relative URI",
		Source:        "RFC 5280: 4.2.1.6",
		EffectiveDate: util.RFC5280Date,
		Test:          &extSANURIRelative{},
	})
}
