// lint_subject_info_access_marked_critical.go
/************************************************
The subject information access extension indicates how to access information and services for the subject of the certificate in which the extension appears. When the subject is a CA, information and services may include certificate validation services and CA policy data. When the subject is an end entity, the information describes the type of services offered and how to access them. In this case, the contents of this extension are defined in the protocol specifications for the supported services. This extension may be included in end entity or CA certificates. Conforming CAs MUST mark this extension as non-critical.
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type siaCrit struct{}

func (l *siaCrit) Initialize() error {
	return nil
}

func (l *siaCrit) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectInfoAccessOID)
}

func (l *siaCrit) Execute(c *x509.Certificate) *LintResult {
	sia := util.GetExtFromCert(c, util.SubjectInfoAccessOID)
	if sia.Critical {
		return &LintResult{Status: Error}
	}
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_subject_info_access_marked_critical",
		Description:   "Conforming CAs MUST mark the Subject Info Access extension as non-critical",
		Citation:      "RFC 5280: 4.2.2.2",
		Source:        RFC5280,
		EffectiveDate: util.RFC3280Date,
		Lint:          &siaCrit{},
	})
}
