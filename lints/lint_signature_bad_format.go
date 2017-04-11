// lint_signature_bad_format.go

package lints

import (
	"bytes"
	"encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SignatureBadFormat struct {
	// Internal data here
}

func (l *SignatureBadFormat) Initialize() error {
	return nil
}

func (l *SignatureBadFormat) CheckApplies(c *x509.Certificate) bool {
	return c.PublicKeyAlgorithm == x509.DSA
}

func (l *SignatureBadFormat) RunTest(c *x509.Certificate) (ResultStruct, error) {
	var seq asn1.RawValue
	// 2nd level sequence with version and serial number
	_, err := asn1.Unmarshal(c.RawTBSCertificate, &seq)
	if err != nil {
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
	// unmarshals signature algorithm identifier into seq
	_, err = asn1.Unmarshal(rest, &seq)
	if err != nil {
		return ResultStruct{Result: NA}, err
	}
	sig := seq.FullBytes
	// unmarshal raw certificate, 1st level sequence
	_, err = asn1.Unmarshal(c.Raw, &seq)
	// unmarshal TBSCertificate
	rest, err = asn1.Unmarshal(seq.Bytes, &seq)
	// unmarshal CA signature
	_, err = asn1.Unmarshal(rest, &seq)
	if bytes.Compare(sig, seq.FullBytes) != 0 {
		return ResultStruct{Result: Error}, nil
	}
	// unmarshals signature algorithm identifier into seq
	_, err = asn1.Unmarshal(rest, &seq)
	if err != nil {
		return ResultStruct{Result: NA}, err
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_signature_bad_format",
		Description:   "Certificate's signature algorithm in Certificate Sequence must have the same algorithm identifier as in TBSCertificate Sequence.",
		Providence:    "RFC 5280 4.1.2.3",
		EffectiveDate: util.RFC3280Date,
		Test:          &SignatureBadFormat{}})
}
