package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type subCertSubjectGnOrSnContainsPolicy struct {
	// Internal data here
}

func (l *subCertSubjectGnOrSnContainsPolicy) Initialize() error {
	return nil
}

func (l *subCertSubjectGnOrSnContainsPolicy) CheckApplies(c *x509.Certificate) bool {
	//Check if GivenName or Surname fields are filled out
	return util.IsSubscriberCert(c) && (len(c.Subject.GivenName) != 0 || len(c.Subject.Surname) != 0)
}

func (l *subCertSubjectGnOrSnContainsPolicy) RunTest(c *x509.Certificate) (ResultStruct, error) {
	for _, policyIds := range c.PolicyIdentifiers {
		if policyIds.Equal(util.BRIndividualValidatedOID) {
			return ResultStruct{Result: Pass}, nil
		}
	}
	return ResultStruct{Result: Error}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_sub_cert_given_name_surname_contains_correct_policy",
		Description:   "Subscriber Certificate: A certificate containing a subject:givenName field or subject:surname field MUST contain the (2.23.140.1.2.3) certPolicy OID.",
		Source:        "BRs: 7.1.4.2.2",
		EffectiveDate: util.CABEffectiveDate,
		Test:          &subCertSubjectGnOrSnContainsPolicy{},
	})
}
