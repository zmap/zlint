package cabf_br

/*
 * ZLint Copyright 2024 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

import (
	"fmt"
	"github.com/zmap/zcrypto/encoding/asn1"
	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

type subjectRdnsCorrectEncoding struct{}

func init() {
	lint.RegisterCertificateLint(&lint.CertificateLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_subject_rdns_correct_encoding",
			Description:   "CAs that include attributes in the Certificate subject field that are listed in the Tables 77 and 78 of BR 2.0.0 SHALL follow the specified encoding requirements for the attribute",
			Citation:      "BRs 2.0.0: 7.1.4.2, Table 77 and Table 78",
			Source:        lint.CABFBaselineRequirements,
			EffectiveDate: util.SC62EffectiveDate,
		},
		Lint: NewSubjectRdnsCorrectEncoding,
	})
}

func NewSubjectRdnsCorrectEncoding() lint.LintInterface {
	return &subjectRdnsCorrectEncoding{}
}

func (l *subjectRdnsCorrectEncoding) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *subjectRdnsCorrectEncoding) Execute(c *x509.Certificate) *lint.LintResult {
	rdnSequence := util.RawRDNSequence{}
	if rest, err := asn1.Unmarshal(c.RawSubject, &rdnSequence); err != nil || len(rest) > 0 {
		return &lint.LintResult{Status: lint.Fatal}
	}

	for _, attrTypeAndValueSet := range rdnSequence {
		for _, attrTypeAndValue := range attrTypeAndValueSet {
			oid := attrTypeAndValue.Type.String()
			tag := attrTypeAndValue.Value.Tag

			if "0.9.2342.19200300.100.1.25" == oid && tag != 22 {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Attribute domainComponent in subjectDN has the wrong encoding %s.", getEncodingName(tag))}
			}
			if "2.5.4.6" == oid && tag != 19 {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Attribute countryName in subjectDN has the wrong encoding %s.", getEncodingName(tag))}
			}
			if "2.5.4.8" == oid && tag != 19 && tag != 12 {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Attribute stateOrProvinceName in subjectDN has the wrong encoding %s.", getEncodingName(tag))}
			}
			if "2.5.4.7" == oid && tag != 19 && tag != 12 {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Attribute localityName in subjectDN has the wrong encoding %s.", getEncodingName(tag))}
			}
			if "2.5.4.17" == oid && tag != 19 && tag != 12 {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Attribute postalCode in subjectDN has the wrong encoding %s.", getEncodingName(tag))}
			}
			if "2.5.4.9" == oid && tag != 19 && tag != 12 {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Attribute streetAddress in subjectDN has the wrong encoding %s.", getEncodingName(tag))}
			}
			if "2.5.4.10" == oid && tag != 19 && tag != 12 {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Attribute organizationName in subjectDN has the wrong encoding %s.", getEncodingName(tag))}
			}
			if "2.5.4.4" == oid && tag != 19 && tag != 12 {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Attribute surname in subjectDN has the wrong encoding %s.", getEncodingName(tag))}
			}
			if "2.5.4.42" == oid && tag != 19 && tag != 12 {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Attribute givenName in subjectDN has the wrong encoding %s.", getEncodingName(tag))}
			}
			if "2.5.4.11" == oid && tag != 19 && tag != 12 {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Attribute organizationalUnitName in subjectDN has the wrong encoding %s.", getEncodingName(tag))}
			}
			if "2.5.4.3" == oid && tag != 19 && tag != 12 {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Attribute commonName in subjectDN has the wrong encoding %s.", getEncodingName(tag))}
			}
			if "2.5.4.15" == oid && tag != 19 && tag != 12 {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Attribute businessCategory in subjectDN has the wrong encoding %s.", getEncodingName(tag))}
			}
			if "1.3.6.1.4.1.311.60.2.1.3" == oid && tag != 19 {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Attribute jurisdictionCountry in subjectDN has the wrong encoding %s.", getEncodingName(tag))}
			}
			if "1.3.6.1.4.1.311.60.2.1.2" == oid && tag != 19 && tag != 12 {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Attribute jurisdictionStateOrProvince in subjectDN has the wrong encoding %s.", getEncodingName(tag))}
			}
			if "1.3.6.1.4.1.311.60.2.1.1" == oid && tag != 19 && tag != 12 {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Attribute jurisdictionLocality in subjectDN has the wrong encoding %s.", getEncodingName(tag))}
			}
			if "2.5.4.5" == oid && tag != 19 {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Attribute serialNumber in subjectDN has the wrong encoding %s.", getEncodingName(tag))}
			}
			if "2.5.4.97" == oid && tag != 19 && tag != 12 {
				return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("Attribute organizationIdentifier in subjectDN has the wrong encoding %s.", getEncodingName(tag))}
			}
		}
	}
	return &lint.LintResult{Status: lint.Pass}
}

//Tag BMPString: 0x1e = 30
//Tag UTF8String: 0x0c = 12
//Tag TeletexString: 0x14 = 20
//Tag UniversalString: 0x1c = 28
//Tag PrintableString: 0x13 = 19
//Tag IA5String: 0x16 = 22

func getEncodingName(tag int) string {
	if tag == 12 {
		return "UTF8String"
	}
	if tag == 19 {
		return "PrintableString"
	}
	if tag == 20 {
		return "TeletexString"
	}
	if tag == 22 {
		return "IA5String"
	}
	if tag == 28 {
		return "UniversalString"
	}
	if tag == 30 {
		return "BMPString"
	}
	return "Unknown"
}
