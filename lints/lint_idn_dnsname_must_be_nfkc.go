package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"golang.org/x/net/idna"
	"golang.org/x/text/unicode/norm"
	"strings"
)

type IDNNotNFKC struct {
	// Internal data here
}

func (l *IDNNotNFKC) Initialize() error {
	return nil
}

func (l *IDNNotNFKC) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *IDNNotNFKC) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dns := range c.DNSNames {
		labels := strings.Split(dns, ".")
		for _, label := range labels {
			if strings.HasPrefix(label, "xn--") {
				unicodeLabel, err := idna.ToUnicode(label)
				if err != nil {
					return ResultStruct{Result: NA}, nil
				}
				if !norm.NFKC.IsNormalString(unicodeLabel) {
					return ResultStruct{Result: Error}, nil
				}
			}
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_international_dns_name_not_nfkc",
		Description:   "Internationalized DNSNames must be normalized by unicode normalization form KC",
		Source:        "RFC 3490",
		EffectiveDate: util.RFC3490Date,
		Test:          &IDNNotNFKC{},
	})
}
