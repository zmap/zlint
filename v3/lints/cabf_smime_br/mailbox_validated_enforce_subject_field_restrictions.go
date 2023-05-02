package cabf_smime_br

import (
	"fmt"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

// MailboxValidatedEnforceSubjectFieldRestrictions - linter to enforce MAY/SHALL NOT requirements for mailbox validated SMIME certificates
type MailboxValidatedEnforceSubjectFieldRestrictions struct{}

var forbiddenSubjectFields = map[string]string{
	"0.9.2342.19200300.100.1.25": "subject:domainComponent",
	"1.3.6.1.4.1.311.60.2.1.1":   "subject:jurisdictionLocality",
	"1.3.6.1.4.1.311.60.2.1.2":   "subject:jurisdictionProvince",
	"1.3.6.1.4.1.311.60.2.1.3":   "subject:jurisdictionCountry",
	"2.5.4.4":                    "subject:surname",
	"2.5.4.6":                    "subject:countryName",
	"2.5.4.7":                    "subject:localityName",
	"2.5.4.8":                    "subject:stateOrProvinceName",
	"2.5.4.9":                    "subject:streetAddress",
	"2.5.4.10":                   "subject:organizationName",
	"2.5.4.11":                   "subject:organizationalUnitName",
	"2.5.4.12":                   "subject:title",
	"2.5.4.17":                   "subject:postalCode",
	"2.5.4.42":                   "subject:givenName",
	"2.5.4.65":                   "subject:pseudonym",
	"2.5.4.97":                   "subject:organizationIdentifier",
}

var acceptableSubjectFields = map[string]string{
	"1.2.840.113549.1.9.1": "subject:emailAddress",
	"2.5.4.3":              "subject:commonName",
	"2.5.4.5":              "subject:serialNumber",
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_mailbox_validated_enforce_subject_field_restrictions",
		Description:   "SMIME certificates complying to mailbox validated profiles MAY only contain commonName, serialNumber or emailAddress attributes in the Subject DN",
		Citation:      "SMIME BRs: 7.1.4.2.3",
		Source:        lint.CABFSMIMEBaselineRequirements,
		EffectiveDate: util.CABF_SMIME_BRs_1_0_0_Date,
		Lint: func() lint.LintInterface {
			return &MailboxValidatedEnforceSubjectFieldRestrictions{}
		},
	})
}

// NewMailboxValidatedEnforceSubjectFieldRestrictions creates a new linter to enforce MAY/SHALL NOT field requirements for mailbox validated SMIME certs
func NewMailboxValidatedEnforceSubjectFieldRestrictions() lint.LintInterface {
	return &MailboxValidatedEnforceSubjectFieldRestrictions{}
}

// CheckApplies is returns true if the certificate's policies assert that it conforms to the mailbox validated SMIME BRs
func (l *MailboxValidatedEnforceSubjectFieldRestrictions) CheckApplies(c *x509.Certificate) bool {
	return util.IsMailboxValidatedCertificate(c)
}

// Execute applies the requirements on what fields are allowed for mailbox validated SMIME certificates
func (l *MailboxValidatedEnforceSubjectFieldRestrictions) Execute(c *x509.Certificate) *lint.LintResult {
	for _, rdnSeq := range c.Subject.OriginalRDNS {
		for _, field := range rdnSeq {
			oidStr := field.Type.String()

			if _, ok := acceptableSubjectFields[oidStr]; !ok {
				if fieldName, knownField := forbiddenSubjectFields[oidStr]; knownField {
					return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("subject DN contains forbidden field: %s (%s)", fieldName, oidStr)}
				}
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("subject DN contains forbidden field: %s", oidStr)}
			}
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}
