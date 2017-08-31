package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"golang.org/x/net/idna"
	"strings"
)

type IDNMalformedUnicode struct {
	// Internal data here
}

func (l *IDNMalformedUnicode) Initialize() error {
	return nil
}

func (l *IDNMalformedUnicode) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *IDNMalformedUnicode) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.DNSNames {
		labels := strings.Split(dns, ".")
		for _, label := range labels {
			if strings.HasPrefix(label, "xn--") {
				_, err := idna.ToUnicode(label)
				if err != nil {
					return ResultStruct{Result: Error}, nil
				}
			}
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_international_dns_name_not_unicode",
		Description:   "Internationalized DNSNames punycode not valid unicode",
		Source:        "RFC 3490",
		EffectiveDate: util.RFC3490Date,
		Test:          &IDNMalformedUnicode{},
	})
}
