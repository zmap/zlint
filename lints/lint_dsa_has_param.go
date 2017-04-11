// lint_dsa_has_param.go

package lints

import (
	"encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type DSAHasParam struct {
	// Internal data here
}

func (l *DSAHasParam) Initialize() error {
	return nil
}

func (l *DSAHasParam) CheckApplies(c *x509.Certificate) bool {
	return c.PublicKeyAlgorithm == x509.DSA
}

func (l *DSAHasParam) RunTest(c *x509.Certificate) (ResultStruct, error) {
	var seq asn1.RawValue
	// 2nd level sequence with version and serial number
	_, err := asn1.Unmarshal(c.RawTBSCertificate, &seq)
	if err != nil {
		return ResultStruct{Result: NA}, err
	}
	if !seq.IsCompound || seq.Tag != asn1.TagSequence || seq.Class != asn1.ClassUniversal {
		err = asn1.StructuralError{Msg: "bad asn1 sequence"}
		return ResultStruct{Result: NA}, err
	}
	rest, err := asn1.Unmarshal(seq.Bytes, &seq)
	if err != nil {
		return ResultStruct{Result: NA}, err
	}
	// optional version present
	if seq.Tag == 0 {
		rest, err = asn1.Unmarshal(rest, &seq)
		if err != nil {
			return ResultStruct{Result: NA}, err
		}
	} else if seq.Tag != 2 {
		err = asn1.StructuralError{Msg: "bad asn1 sequence"}
		return ResultStruct{Result: NA}, err
	}
	rest, err = asn1.Unmarshal(rest, &seq)
	if err != nil {
		return ResultStruct{Result: NA}, err
	}
	// get signature algorithm oid
	rest, err = asn1.Unmarshal(seq.Bytes, &seq)
	if err != nil {
		return ResultStruct{Result: NA}, err
	}
	if len(rest) > 0 {
		return ResultStruct{Result: Error}, nil
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_dsa_has_param",
		Description:   "DSA keys must not have a parameter specified",
		Providence:    "Certlint",
		EffectiveDate: util.ZeroDate,
		Test:          &DSAHasParam{}})
}
