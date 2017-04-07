// lint_issuer_dn_trailing_whitespace.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type IssuerDNTrailingSpace struct {
	// Internal data here
}

func (l *IssuerDNTrailingSpace) Initialize() error {
	return nil
}

func (l *IssuerDNTrailingSpace) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *IssuerDNTrailingSpace) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, country := range c.Issuer.Country {
		if strings.HasSuffix(country, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, org := range c.Issuer.Organization{
		if strings.HasSuffix(org, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, unit := range c.Issuer.OrganizationalUnit{
		if strings.HasSuffix(unit, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, unit := range c.Issuer.OrganizationalUnit{
		if strings.HasSuffix(unit, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, loc := range c.Issuer.Locality{
		if strings.HasSuffix(loc, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, prov := range c.Issuer.Province{
		if strings.HasSuffix(prov, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, addr := range c.Issuer.StreetAddress{
		if strings.HasSuffix(addr, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, post := range c.Issuer.PostalCode{
		if strings.HasSuffix(post, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, domain := range c.Issuer.DomainComponent{
		if strings.HasSuffix(domain, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	if strings.HasSuffix(c.Issuer.SerialNumber, " ") {
		return ResultStruct{Result: Warn}, nil
	}
	if strings.HasSuffix(c.Issuer.CommonName, " ") {
		return ResultStruct{Result: Warn}, nil
	}
	return ResultStruct{Result: Pass}, nil
}
	
func init() {
	RegisterLint(&Lint{
		Name:          "w_issuer_dn_trailing_whitespace",
		Description:   "Issuer distinguished name attribute should not have trailing whitespace",
		Providence:    "Certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &IssuerDNTrailingSpace{}})
}
