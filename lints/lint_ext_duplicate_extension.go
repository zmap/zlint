/************************************************
"A certificate MUST NOT include more than one instance of a particular extension."
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type ExtDuplicateExtension struct{}

func (l *ExtDuplicateExtension) Initialize() error {
	return nil
}

func (l *ExtDuplicateExtension) CheckApplies(cert *x509.Certificate) bool {
	return cert.Version == 3
}

func (l *ExtDuplicateExtension) Execute(cert *x509.Certificate) *LintResult {
	// O(n^2) is not terrible here because n is capped around 10
	for i := 0; i < len(cert.Extensions); i++ {
		for j := i + 1; j < len(cert.Extensions); j++ {
			if i != j && cert.Extensions[i].Id.Equal(cert.Extensions[j].Id) {
				return &LintResult{Status: Error}
			}
		}
	}
	// Nested loop will return if it finds a duplicate, so safe to assume pass
	return &LintResult{Status: Pass}
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_duplicate_extension",
		Description:   "A certificate MUST NOT include more than one instance of a particular extension",
		Citation:      "RFC 5280: 4.2",
		Source:        RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          &ExtDuplicateExtension{},
	})
}
