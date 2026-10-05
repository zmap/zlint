/*
 * ZLint Copyright 2026 Regents of the University of Michigan
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

package cabf_br

import (
	"encoding/asn1"
	"errors"
	"fmt"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/util"
)

func init() {
	lint.RegisterRevocationListLint(&lint.RevocationListLint{
		LintMetadata: lint.LintMetadata{
			Name:          "e_crl_next_update_invalid",
			Description:   "For CRLs covering (EE|CA) certificates, nextUpdate must be at most (10 days|12 months) beyond thisUpdate",
			Citation:      "Section 4.9.7 of BRs v1.8.7 (then section 7.2 since BRs v2.0.0)",
			Source:        lint.CABFBaselineRequirements,
			EffectiveDate: util.CABFBRs_1_8_7_Date,
		},
		Lint: NewCrlNextUpdateInvalid,
	})
}

type CrlNextUpdateInvalid struct {
	SubscriberCRL bool `comment:"Set this to false if the CRL to be linted covers CA certificates"`
}

func (l *CrlNextUpdateInvalid) Configure() interface{} {
	return l
}

func NewCrlNextUpdateInvalid() lint.RevocationListLintInterface {
	return &CrlNextUpdateInvalid{
		SubscriberCRL: true,
	}
}

func (l *CrlNextUpdateInvalid) CheckApplies(c *x509.RevocationList) bool {
	// If NextUpdate is absent it's an error but it's not this lint's business
	return !c.NextUpdate.IsZero()
}

func (l *CrlNextUpdateInvalid) Execute(c *x509.RevocationList) *lint.LintResult {

	// As set out in the CABF BRs
	CabfMaxEECRLValidityDays := 10
	CabfMaxCACRLValidityMonths := 12

	// Which limit applies depends on whether the CRL covers subscriber or CA
	// certificates, and that cannot be determined from an arbitrary CRL. The
	// SubscriberCRL configuration field is therefore used by default.
	//
	// However, when the CRL carries an Issuing Distribution Point (RFC 5280,
	// Section 5.2.5) that explicitly scopes itself to CA or subscriber
	// certificates, that unambiguous in-band signal is preferred over the
	// configuration default.
	subscriberCRL := l.SubscriberCRL
	switch scope, err := idpCRLScope(c); {
	case err != nil:
		return &lint.LintResult{Status: lint.Error, Details: err.Error()}
	case scope == crlScopeCA:
		subscriberCRL = false
	case scope == crlScopeSubscriber:
		subscriberCRL = true
	}

	if subscriberCRL {
		if c.NextUpdate.After(c.ThisUpdate.AddDate(0, 0, CabfMaxEECRLValidityDays)) {
			return &lint.LintResult{
				Status: lint.Error,
				Details: fmt.Sprintf(
					"For CRLs covering Subscriber Certificates, nextUpdate must be at most %d days after thisUpdate",
					CabfMaxEECRLValidityDays),
			}
		}
	} else {
		if c.NextUpdate.After(c.ThisUpdate.AddDate(0, CabfMaxCACRLValidityMonths, 0)) {
			return &lint.LintResult{
				Status: lint.Error,
				Details: fmt.Sprintf(
					"For CRLs covering CA Certificates, nextUpdate must be at most %d months after thisUpdate",
					CabfMaxCACRLValidityMonths),
			}
		}
	}

	return &lint.LintResult{Status: lint.Pass}
}

var errAmbiguousIDPScope = errors.New("IDP does not unambiguously scope the CRL")

//	IssuingDistributionPoint ::= SEQUENCE {
//	     distributionPoint          [0] DistributionPointName OPTIONAL,
//	     onlyContainsUserCerts      [1] BOOLEAN DEFAULT FALSE,
//	     onlyContainsCACerts        [2] BOOLEAN DEFAULT FALSE,
//	     onlySomeReasons            [3] ReasonFlags OPTIONAL,
//	     indirectCRL                [4] BOOLEAN DEFAULT FALSE,
//	     onlyContainsAttributeCerts [5] BOOLEAN DEFAULT FALSE }
type issuingDistPoint struct {
	DistributionPoint     asn1.RawValue  `asn1:"optional,tag:0"`
	OnlyContainsUserCerts bool           `asn1:"optional,tag:1"`
	OnlyContainsCACerts   bool           `asn1:"optional,tag:2"`
	OnlySomeReasons       asn1.BitString `asn1:"optional,tag:3"`
	IndirectCRL           bool           `asn1:"optional,tag:4"`
	OnlyContainsAttrCerts bool           `asn1:"optional,tag:5"`
}

type crlScope int

const (
	crlScopeUnknown crlScope = iota
	crlScopeSubscriber
	crlScopeCA
)

// idpCRLScope inspects the CRL's Issuing Distribution Point extension, if
// present, and reports whether it unambiguously scopes the CRL to CA or
// subscriber certificates. It returns (crlScopeUnknown, nil) when there is
// no IDP extension, or when it's present but doesn't assert either scope
// boolean. It returns a non-nil error when the extension can't be parsed,
// or when it asserts both scope booleans (which BR §7.2.2.1 prohibits).
func idpCRLScope(c *x509.RevocationList) (crlScope, error) {
	for _, ext := range c.Extensions {
		if !ext.Id.Equal(util.IssuingDistOID) {
			continue
		}

		var idp issuingDistPoint
		if _, err := asn1.Unmarshal(ext.Value, &idp); err != nil {
			return crlScopeUnknown, err
		}

		switch {
		case idp.OnlyContainsCACerts && idp.OnlyContainsUserCerts:
			return crlScopeUnknown, errAmbiguousIDPScope
		case idp.OnlyContainsCACerts:
			return crlScopeCA, nil
		case idp.OnlyContainsUserCerts:
			return crlScopeSubscriber, nil
		default:
			return crlScopeUnknown, nil
		}
	}
	return crlScopeUnknown, nil
}
