// lint_rsa_not_null_param.go
/********************************************************************
The rsaEncryption OID is intended to be used in the algorithm field
of a value of type AlgorithmIdentifier.  The parameters field MUST
have ASN.1 type NULL for this algorithm identifier.
********************************************************************/

package lints

import (
	"encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type RSANotNullParam struct {
	// Internal data here
}

func (l *RSANotNullParam) Initialize() error {
	return nil
}

func (l *RSANotNullParam) CheckApplies(c *x509.Certificate) bool {
	return c.PublicKeyAlgorithm == x509.RSA
}

func (l *RSANotNullParam) RunTest(c *x509.Certificate) (ResultStruct, error) {
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
	if  seq.Tag == 0 {
		rest, err = asn1.Unmarshal(rest, &seq)
		if err != nil {
			return ResultStruct{Result: NA}, err
		}
	} else if seq.Tag != 2 {
		err = asn1.StructuralError{Msg: "bad asn1 sequence"}
		return ResultStruct{Result: NA}, err
	}
	// skip through 5 sequences to get to get to SPKI struct
	for i := 0; i < 5; i++ {
		rest, err = asn1.Unmarshal(rest, &seq)
		if err != nil {
			return ResultStruct{Result: NA}, err
		}
	}
	// get public key sequence
	_, err = asn1.Unmarshal(seq.Bytes, &seq)
	if err != nil {
		return ResultStruct{Result: NA}, err
	}

	// public key algorithm in seq
	rest, err = asn1.Unmarshal(seq.Bytes, &seq)
	if err != nil {
		return ResultStruct{Result: NA}, err
	}

	if len(rest) == 0 {
		return ResultStruct{Result: NA}, err
	}
	// public key parameter
	rest, err = asn1.Unmarshal(rest, &seq)
	if err != nil {
		return ResultStruct{Result: NA}, err
	}
	if seq.Tag != 5 {
		return ResultStruct{Result: Error}, nil
	}
	return ResultStruct{Result: Pass}, nil
}
	
func init() {
	RegisterLint(&Lint{
		Name:          "e_rsa_not_null_param",
		Description:   "RSA keys must have a parameter specified",
		Providence:    "RFC 3279 2.3.1",
		EffectiveDate: util.RFC3280Date,
		Test:          &RSANotNullParam{}})
}
