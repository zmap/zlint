// lint_ext_san_uri_format_invalid.go
/************************************************
The name MUST include both a
scheme (e.g., "http" or "ftp") and a scheme-specific-part.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"net/url"
)

type extSANURIFormatInvalid struct {
	// Internal data here
}

func (l *extSANURIFormatInvalid) Initialize() error {
	return nil
}

func (l *extSANURIFormatInvalid) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *extSANURIFormatInvalid) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, uri := range c.URIs {
		parsed_uri, err := url.Parse(uri)

		if err != nil {
			return ResultStruct{Result: Error}, nil
		}

		//scheme
		if parsed_uri.Scheme == "" {
			return ResultStruct{Result: Error}, nil
		}

		//scheme-specific part
		if parsed_uri.Host == "" && parsed_uri.User == nil && parsed_uri.Opaque == "" && parsed_uri.Path == "" {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_uri_format_invalid",
		Description:   "URIs in SAN extension must have a scheme and scheme specific part",
		Source:        "RFC5280: 4.2.1.6",
		EffectiveDate: util.RFC5280Date,
		Test:          &extSANURIFormatInvalid{},
	})
}
