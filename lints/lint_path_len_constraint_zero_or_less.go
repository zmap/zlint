// lint_path_len_constraint_zero_or_less.go
/********************************************************************
The pathLenConstraint field is meaningful only if the cA boolean is
asserted and the key usage extension, if present, asserts the
keyCertSign bit (Section 4.2.1.3).  In this case, it gives the
maximum number of non-self-issued intermediate certificates that may
follow this certificate in a valid certification path.  (Note: The
last certificate in the certification path is not an intermediate
certificate, and is not included in this limit.  Usually, the last
certificate is an end entity certificate, but it can be a CA
certificate.)  A pathLenConstraint of zero indicates that no non-
self-issued intermediate CA certificates may follow in a valid
certification path.  Where it appears, the pathLenConstraint field
MUST be greater than or equal to zero.  Where pathLenConstraint does
not appear, no limit is imposed.
********************************************************************/

package lints

import (
	"encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type basicConst struct {
	CA                bool `asn1:"optional"`
	PathLenConstraint int  `asn1:"optional"`
}

type pathLenNonPositive struct {
	bc basicConst
}

func (l *pathLenNonPositive) Initialize() error {
	return nil
}

func (l *pathLenNonPositive) CheckApplies(cert *x509.Certificate) bool {
	return cert.BasicConstraintsValid
}

func (l *pathLenNonPositive) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	ext := util.GetExtFromCert(cert, util.BasicConstOID)
	_, err := asn1.Unmarshal(ext.Value, &l.bc)
	if err != nil {
		return ResultStruct{Result: Fatal}, err
	}
	if l.bc.PathLenConstraint < 0 {
		return ResultStruct{Result: Error}, err
	}
	return ResultStruct{Result: Pass}, err
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_path_len_constraint_zero_or_less",
		Description:   "Where it appears, the pathLenConstraint field MUST be greater than or equal to zero",
		Source:        "RFC 5280: 4.2.1.9",
		EffectiveDate: util.RFC2459Date,
		Test:          &pathLenNonPositive{},
	})
}
