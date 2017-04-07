// lint_subject_dn_leading_whitespace.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type SubjectDNLeadingSpace struct {
	// Internal data here
}

func (l *SubjectDNLeadingSpace) Initialize() error {
	return nil
}

func (l *SubjectDNLeadingSpace) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *SubjectDNLeadingSpace) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, country := range c.Subject.Country {
		if strings.HasPrefix(country, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, org := range c.Subject.Organization{
		if strings.HasPrefix(org, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, unit := range c.Subject.OrganizationalUnit{
		if strings.HasPrefix(unit, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, unit := range c.Subject.OrganizationalUnit{
		if strings.HasPrefix(unit, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, loc := range c.Subject.Locality{
		if strings.HasPrefix(loc, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, prov := range c.Subject.Province{
		if strings.HasPrefix(prov, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, addr := range c.Subject.StreetAddress{
		if strings.HasPrefix(addr, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, post := range c.Subject.PostalCode{
		if strings.HasPrefix(post, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, domain := range c.Subject.DomainComponent{
		if strings.HasPrefix(domain, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	if strings.HasPrefix(c.Subject.SerialNumber, " ") {
		return ResultStruct{Result: Warn}, nil
	}
	if strings.HasPrefix(c.Subject.CommonName, " ") {
		return ResultStruct{Result: Warn}, nil
	}
	return ResultStruct{Result: Pass}, nil
}
	
func init() {
	RegisterLint(&Lint{
		Name:          "w_subject_dn_leading_whitespace",
		Description:   "Subject distinguished name attribute should not have leading whitespace",
		Providence:    "Certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &SubjectDNLeadingSpace{}})
}
