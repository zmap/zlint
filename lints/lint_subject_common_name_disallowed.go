// lint_subject_common_name_disallowed.go
/************************************************************************************************
7.1.4.2.2. Subject Distinguished Name Fields
a. Certificate Field: subject:commonName (OID 2.5.4.3)
Required/Optional: Deprecated (Discouraged, but not prohibited)
Contents: If present, this field MUST contain a single IP address or Fully‐Qualified Domain
Name that is one of the values contained in the Certificate’s subjectAltName extension (see
Section 7.1.4.2.1).
************************************************************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type BadCommonName struct{}

func (l *BadCommonName) Initialize() error {
	return nil
}

func (l *BadCommonName) CheckApplies(c *x509.Certificate) bool {
	return len(c.Subject.CommonName) != 0 && !util.IsCACert(c)
}

func (l *BadCommonName) Execute(c *x509.Certificate) * LintResult{
	for _, dns := range c.DNSNames {
		if dns == c.Subject.CommonName {
			return &LintResult{Status: Pass}
		}
	}
	for _, ip := range c.IPAddresses {
		if ip.String() == c.Subject.CommonName {
			return &LintResult{Status: Pass}
		}
	}
	return &LintResult{Status: Error}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_common_name_disallowed",
		Description:   "If present in a subscriber certificate, commonName MUST contain a single IP address or Fully‐Qualified Domain Name that is one of the values contained in the certificate’s subjectAltName extension",
		Source:        "CAB: 7.1.4.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Lint:          &BadCommonName{},
	})
}
