// lint_ub_common_name_invalid.go
/************************************************
RFC 5280: A.1
	* In this Appendix, there is a list of upperbounds
	for fields in a x509 Certificate. *
	ub-common-name INTEGER ::= 64
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type ubCommonNameInvalid struct {
	// Internal data here
}

func (l *ubCommonNameInvalid) Initialize() error {
	return nil
}

func (l *ubCommonNameInvalid) CheckApplies(c *x509.Certificate) bool {
	// Add conditions for application here
	return true
}

func (l *ubCommonNameInvalid) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if len(c.Subject.CommonName) <= 64 {
		return ResultStruct{Result: Pass}, nil
	} else {
		return ResultStruct{Result: Error}, nil
	}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ub_common_name_invalid",
		Description:   "The 'Common Name' field must be less than 64 integers long.",
		Providence:    "RFC 5280: A.1",
		EffectiveDate: util.RFC2459Date,
		Test:          &ubCommonNameInvalid{}})
}
