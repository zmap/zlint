// lint_issuer_dn_leading_whitespace.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type IssuerDNLeadingSpace struct {
	// Internal data here
}

func (l *IssuerDNLeadingSpace) Initialize() error {
	return nil
}

func (l *IssuerDNLeadingSpace) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *IssuerDNLeadingSpace) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, country := range c.Issuer.Country {
		if strings.HasPrefix(country, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, org := range c.Issuer.Organization{
		if strings.HasPrefix(org, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, unit := range c.Issuer.OrganizationalUnit{
		if strings.HasPrefix(unit, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, unit := range c.Issuer.OrganizationalUnit{
		if strings.HasPrefix(unit, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, loc := range c.Issuer.Locality{
		if strings.HasPrefix(loc, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, prov := range c.Issuer.Province{
		if strings.HasPrefix(prov, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, addr := range c.Issuer.StreetAddress{
		if strings.HasPrefix(addr, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, post := range c.Issuer.PostalCode{
		if strings.HasPrefix(post, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, domain := range c.Issuer.DomainComponent{
		if strings.HasPrefix(domain, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	if strings.HasPrefix(c.Issuer.SerialNumber, " ") {
		return ResultStruct{Result: Warn}, nil
	}
	if strings.HasPrefix(c.Issuer.CommonName, " ") {
		return ResultStruct{Result: Warn}, nil
	}
	return ResultStruct{Result: Pass}, nil
}
	
func init() {
	RegisterLint(&Lint{
		Name:          "w_issuer_dn_leading_whitespace",
		Description:   "Issuer distinguished name attribute should not have leading whitespace",
		Providence:    "Certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &IssuerDNLeadingSpace{}})
}
