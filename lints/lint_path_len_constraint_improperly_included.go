// lint_path_len_constraint_improperly_included.go
/******************************************************************
RFC 5280: 4.2.1.9
CAs MUST NOT include the pathLenConstraint field unless the cA
boolean is asserted and the key usage extension asserts the
keyCertSign bit.
******************************************************************/

package lints

import (
	"encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type pathLenIncluded struct {
	// Internal data here
}

func (l *pathLenIncluded) Initialize() error {
	return nil
}

func (l *pathLenIncluded) CheckApplies(cert *x509.Certificate) bool {
	return util.IsExtInCert(cert, util.BasicConstOID)
}

func (l *pathLenIncluded) RunTest(cert *x509.Certificate) (ResultStruct, error) {
	bc := util.GetExtFromCert(cert, util.BasicConstOID)
	var seq asn1.RawValue
	var isCa bool
	_, err := asn1.Unmarshal(bc.Value, &seq)
	if err != nil {
		return ResultStruct{Result: Fatal}, nil
	}
	if len(seq.Bytes) == 0 {
		return ResultStruct{Result: Pass}, nil
	}
	rest, err := asn1.UnmarshalWithParams(seq.Bytes, &isCa, "optional")
	if err != nil {
		return ResultStruct{Result: Fatal}, err
	}
	kUVal := util.IsExtInCert(cert, util.KeyUsageOID)
	if len(rest) > 0 && (!cert.IsCA || !kUVal || (kUVal && cert.KeyUsage&x509.KeyUsageCertSign == 0)) {
		return ResultStruct{Result: Error}, nil
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_path_len_constraint_improperly_included",
		Description:   "CAs MUST NOT include the pathLenConstraint field unless the CA boolean is asserted and the keyCertSign bit is set",
		Source:        "RFC 5280: 4.2.1.9",
		EffectiveDate: util.RFC3280Date,
		Test:          &pathLenIncluded{},
	})
}
