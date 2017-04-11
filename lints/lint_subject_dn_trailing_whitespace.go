// lint_subject_dn_trailing_whitespace.go

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
	"strings"
)

type SubjectDNTrailingSpace struct {
	// Internal data here
}

func (l *SubjectDNTrailingSpace) Initialize() error {
	return nil
}

func (l *SubjectDNTrailingSpace) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *SubjectDNTrailingSpace) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, country := range c.Subject.Country {
		if strings.HasSuffix(country, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, org := range c.Subject.Organization{
		if strings.HasSuffix(org, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, unit := range c.Subject.OrganizationalUnit{
		if strings.HasSuffix(unit, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, loc := range c.Subject.Locality{
		if strings.HasSuffix(loc, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, prov := range c.Subject.Province{
		if strings.HasSuffix(prov, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, addr := range c.Subject.StreetAddress{
		if strings.HasSuffix(addr, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, post := range c.Subject.PostalCode{
		if strings.HasSuffix(post, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	for _, domain := range c.Subject.DomainComponent{
		if strings.HasSuffix(domain, " ") {
			return ResultStruct{Result: Warn}, nil
		}
	}
	if strings.HasSuffix(c.Subject.SerialNumber, " ") {
		return ResultStruct{Result: Warn}, nil
	}
	if strings.HasSuffix(c.Subject.CommonName, " ") {
		return ResultStruct{Result: Warn}, nil
	}
	return ResultStruct{Result: Pass}, nil
}
	
func init() {
	RegisterLint(&Lint{
		Name:          "w_subject_dn_trailing_whitespace",
		Description:   "Subject distinguished name attribute should not have trailing whitespace",
		Providence:    "Certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &SubjectDNTrailingSpace{}})
}
