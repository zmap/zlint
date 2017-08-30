// lint_ext_duplicate_extension.go
/************************************************
"A certificate MUST NOT include more than one instance of a particular extension."
************************************************/

package lints

import (
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type ExtDuplicateExtension struct {
	// Internal data here
}

func (l *ExtDuplicateExtension) Initialize() error {
	return nil
}

func (l *ExtDuplicateExtension) CheckApplies(cert *x509.Certificate) bool {
	return cert.Version == 3
}

func (l *ExtDuplicateExtension) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	// O(n^2) is not terrible here because n is capped around 10
	for i := 0; i < len(cert.Extensions); i++ {
		for j := i + 1; j < len(cert.Extensions); j++ {
			if i != j && cert.Extensions[i].Id.Equal(cert.Extensions[j].Id) {
				return ResultStruct{Result: Error}, nil
			}
		}
	}
	// Nested loop will return if it finds a duplicate, so safe to assume pass
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_duplicate_extension",
		Description:   "A certificate MUST NOT include more than one instance of a particular extension",
		Source:        "RFC 5280: 4.2",
		EffectiveDate: util.RFC2459Date,
		Test:          &ExtDuplicateExtension{},
	})
}
