package util

import (
	"encoding/asn1"
	"errors"
	"github.com/zmap/zgrab/ztools/x509"
	"github.com/zmap/zgrab/ztools/x509/pkix"
)

var (
	//extension OIDs
	AiaOID                = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 1, 1}        // Authority Information Access
	AuthkeyOID            = asn1.ObjectIdentifier{2, 5, 29, 35}                     // Authority Key Identifier
	BasicConstOID         = asn1.ObjectIdentifier{2, 5, 29, 19}                     // Basic Constraints
	CertPolicyOID         = asn1.ObjectIdentifier{2, 5, 29, 32}                     // Certificate Policies
	CrlDistOID            = asn1.ObjectIdentifier{2, 5, 29, 31}                     // CRL Distribution Points
	CtPoisonOID           = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 11129, 2, 4, 3} // CT Poison
	EkuSynOid             = asn1.ObjectIdentifier{2, 5, 29, 37}                     // Extended Key Usage Syntax
	FreshCRLOID           = asn1.ObjectIdentifier{2, 5, 29, 46}                     // Freshest CRL
	InhibitAnyPolicyOID   = asn1.ObjectIdentifier{2, 5, 29, 54}                     // Inhibit Any Policy
	IssuerANOID           = asn1.ObjectIdentifier{2, 5, 29, 18}                     // Issuer Alt Name
	KeyUsageOID           = asn1.ObjectIdentifier{2, 5, 29, 15}                     // Key Usage
	LogoTypeOID           = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 1, 12}       // Logo Type Ext
	NameConstOID          = asn1.ObjectIdentifier{2, 5, 29, 30}                     // Name Constraints
	OscpNoCheckOID        = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 48, 1, 5}    // OSCP No Check
	PolicyConstOID        = asn1.ObjectIdentifier{2, 5, 29, 36}                     // Policy Constraints
	PolicyMapOID          = asn1.ObjectIdentifier{2, 5, 29, 33}                     // Policy Mappings
	PrivKeyUsageOID       = asn1.ObjectIdentifier{2, 5, 29, 16}                     // Private Key Usage Period
	QcStateOid            = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 1, 3}        // QC Statements
	TimestampOID          = asn1.ObjectIdentifier{1, 3, 6, 1, 4, 1, 11129, 2, 4, 2} // Signed Certificate Timestamp List
	SmimeOID              = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 15}      // Smime Capabilities
	SanOID                = asn1.ObjectIdentifier{2, 5, 29, 17}                     // Subject Alt Name
	SubjectDirAttrOID     = asn1.ObjectIdentifier{2, 5, 29, 9}                      // Subject Directory Attributes
	SubjectInfoAccessOID  = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 1, 11}       // Subject Info Access Syntax
	SubjectKeyIdentityOID = asn1.ObjectIdentifier{2, 5, 29, 14}                     // Subject Key Identifier
	// CA/B reserved policies
	BRDomainValidatedOID       = asn1.ObjectIdentifier{2, 23, 140, 1, 2, 1} // CA/B BR Domain-Validated
	BROrganizationValidatedOID = asn1.ObjectIdentifier{2, 23, 140, 1, 2, 2} // CA/B BR Organization-Validated
	BRIndividualValidatedOID   = asn1.ObjectIdentifier{2, 23, 140, 1, 2, 3} // CA/B BR Individual-Validated
	//X.500 attribute types
	CommonNameOID             = asn1.ObjectIdentifier{2, 5, 4, 3}
	SurnameOID                = asn1.ObjectIdentifier{2, 5, 4, 4}
	SerialOID                 = asn1.ObjectIdentifier{2, 5, 4, 5}
	CountryNameOID            = asn1.ObjectIdentifier{2, 5, 4, 6}
	LocalityNameOID           = asn1.ObjectIdentifier{2, 5, 4, 7}
	StateOrProvinceNameOID    = asn1.ObjectIdentifier{2, 5, 4, 8}
	StreetAddressOID          = asn1.ObjectIdentifier{2, 5, 4, 9}
	OrganizationNameOID       = asn1.ObjectIdentifier{2, 5, 4, 10}
	OrganizationalUnitNameOID = asn1.ObjectIdentifier{2, 5, 4, 11}
	BusinessOID               = asn1.ObjectIdentifier{2, 5, 4, 15}
	PostalCodeOID             = asn1.ObjectIdentifier{2, 5, 4, 17}
	GivenNameOID              = asn1.ObjectIdentifier{2, 5, 4, 42}
	// other OIDs
	OidRSASSAPSS  = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 10}
	AnyPolicyOID  = asn1.ObjectIdentifier{2, 5, 29, 32, 0}
	UserNoticeOID = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 2, 2}
	CpsOID        = asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 2, 1}
)

func IsExtInCert(cert *x509.Certificate, oid asn1.ObjectIdentifier) bool {
	if cert != nil && GetExtFromCert(cert, oid) != nil {
		return true
	} //else
	return false
}

// Helper function that should be used anytime an extension is needed from a certificate
func GetExtFromCert(cert *x509.Certificate, oid asn1.ObjectIdentifier) *pkix.Extension {
	for i, _ := range cert.Extensions {
		if oid.Equal(cert.Extensions[i].Id) {
			return &(cert.Extensions[i])
		}
	}
	return nil
}

// Helper function that checks if an []asn1.ObjectIdentifier slice contains an asn1.ObjectIdentifier
func SliceContainsOID(list []asn1.ObjectIdentifier, oid asn1.ObjectIdentifier) bool {
	for _, v := range list {
		if oid.Equal(v) {
			return true
		}
	}
	return false
}

// Helper function that checks for a name type in a pkix.Name
func TypeInName(name *pkix.Name, oid asn1.ObjectIdentifier) bool {
	for _, v := range name.Names {
		if oid.Equal(v.Type) {
			return true
		}
	}
	return false
}

//helper function to parse policyMapping extensions, returns slices of CertPolicyIds seperated by domain
func GetMappedPolicies(polMap *pkix.Extension) (out [][2]asn1.ObjectIdentifier, err error) {
	if polMap == nil {
		return nil, errors.New("policyMap: null pointer")
	}
	var outSeq, inSeq asn1.RawValue

	empty, err := asn1.Unmarshal(polMap.Value, &outSeq) //strip outer sequence tag/length should be nothing extra
	if err != nil || len(empty) != 0 || outSeq.Class != 0 || outSeq.Tag != 16 || outSeq.IsCompound == false {
		return nil, errors.New("policyMap: Could not unmarshal outer sequence.")
	}

	for done := false; !done; { //loop through SEQUENCE OF
		outSeq.Bytes, err = asn1.Unmarshal(outSeq.Bytes, &inSeq) //extract next inner SEQUENCE (OID pair)
		if err != nil || inSeq.Class != 0 || inSeq.Tag != 16 || inSeq.IsCompound == false {
			err = errors.New("policyMap: Could not unmarshal inner sequence.")
			return
		}
		if len(outSeq.Bytes) == 0 { //nothing remaining to parse, stop looping after
			done = true
		}

		var oidIssue, oidSubject asn1.ObjectIdentifier
		var restIn asn1.RawContent
		restIn, err = asn1.Unmarshal(inSeq.Bytes, &oidIssue) //extract first inner CertPolicyId (issuer domain)
		if err != nil || len(restIn) == 0 {
			err = errors.New("policyMap: Could not unmarshal inner sequence.")
			return
		}

		empty, err = asn1.Unmarshal(restIn, &oidSubject) //extract second inner CertPolicyId (subject domain)
		if err != nil || len(empty) != 0 {
			err = errors.New("policyMap: Could not unmarshal inner sequence.")
			return
		}

		//append found OIDs
		out = append(out, [2]asn1.ObjectIdentifier{oidIssue, oidSubject})
	}

	return
}
