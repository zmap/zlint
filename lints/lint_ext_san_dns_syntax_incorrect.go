// lint_ext_san_dns_syntax_incorrect.go
/************************************************
RFC 5280: 4.2.1.6
When the subjectAltName extension contains a domain name system
label, the domain name MUST be stored in the dNSName (an IA5String).
The name MUST be in the "preferred name syntax", as specified by
Section 3.5 of [RFC1034] and as modified by Section 2.1 of [RFC1123].
************************************************/

package lints

import (

	"github.com/zmap/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
	"net"
)

type sanDNSPrefSyntax struct {
	// Internal data here
}

func (l *sanDNSPrefSyntax) Initialize() error {
	return nil
}

func (l *sanDNSPrefSyntax) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SanOID)
}

func (l *sanDNSPrefSyntax) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, dnsname := range c.DNSNames {
		// Make sure the dnsname isn't an IP, which auto-passes
		if net.ParseIP(dnsname) != nil {
			continue
		}
		if dnsname != "" && !util.IsInPrefSyn(dnsname) {
			return ResultStruct{Result: Error}, nil
		}
	}

	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "ext_san_dns_syntax_incorrect",
		Description:   "DNSNames must be in the preferred syntax.",
		Providence:    "RFC 5280: 4.2.1.6",
		EffectiveDate: util.RFC3280Date,
		Test:          &sanDNSPrefSyntax{}})
}
