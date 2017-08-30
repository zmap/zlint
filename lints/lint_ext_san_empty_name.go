// lint_ext_san_empty_name.go
/******************************************************************
RFC 5280: 4.2.1.6
If the subjectAltName extension is present, the sequence MUST contain
at least one entry.  Unlike the subject field, conforming CAs MUST
NOT issue certificates with subjectAltNames containing empty
GeneralName fields.  For example, an rfc822Name is represented as an
IA5String.  While an empty string is a valid IA5String, such an
rfc822Name is not permitted by this profile.  The behavior of clients
that encounter such a certificate when processing a certification
path is not defined by this profile.
******************************************************************/

package lints

import (
	"encoding/asn1"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/util"
)

type SANEmptyName struct {
	// Internal data here
}

func (l *SANEmptyName) Initialize() error {
	return nil
}

func (l *SANEmptyName) CheckApplies(c *x509.Certificate) bool {
	return util.IsExtInCert(c, util.SubjectAlternateNameOID)
}

func (l *SANEmptyName) RunTest(c *x509.Certificate) (ResultStruct, error) {
	value := util.GetExtFromCert(c, util.SubjectAlternateNameOID).Value
	var seq asn1.RawValue
	var err error
	if _, err = asn1.Unmarshal(value, &seq); err != nil {
		return ResultStruct{Result: NA}, err
	}
	if !seq.IsCompound || seq.Tag != 16 || seq.Class != 0 {
		err = asn1.StructuralError{Msg: "bad SAN sequence"}
		return ResultStruct{Result: Fatal}, err
	}

	rest := seq.Bytes
	for len(rest) > 0 {

		var v asn1.RawValue
		rest, err = asn1.Unmarshal(rest, &v)
		if err != nil {
			return ResultStruct{Result: NA}, err
		}
		if len(v.Bytes) == 0 {
			return ResultStruct{Result: Error}, nil
		}
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "e_ext_san_empty_name",
		Description:   "General name fields MUST NOT be empty in subjectAlternateNames",
		Source:        "RFC 5280: 4.2.1.6",
		EffectiveDate: util.RFC2459Date,
		Test:          &SANEmptyName{},
	})
}
