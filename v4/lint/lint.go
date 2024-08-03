// ZLint Copyright 2024 Regents of the University of Michigan
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy
// of the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.
//

// Package lint defines the core data structures used for ZLint. New in Version
// 4 is the concept of a Profile.
package lint

import (
	"time"

	"github.com/zmap/zcrypto/x509"
)

// Profile is an enum indicating the certificate profile a certificate should be
// linted against, e.g. Subscriber vs Intermediate vs Root. Lints apply to only
// a subset of profiles.
//

type Profile struct {
	Name    string         `json:"name"`
	Handle  uint64         `json:"-"`
	Matcher ProfileMatcher `json:"-"`
}

type ProfileMatcher interface {
	CheckMatches(c *x509.Certificate) bool
}

func (pd *Profile) CheckMatches(c *x509.Certificate) bool {
	return pd.Matcher.CheckMatches(c)
}

// LintMetadata represents the metadata that are broadly associated across all types of lints.
//
// That is, all lints (irrespective of being a certificate lint, a CRL lint, and OCSP, etc.)
// have a Name, a Description, a Citation, and so on.
//
// In this way, this struct may be embedded in any linting type in order to maintain this
// data, while each individual linting type provides the behavior over this data.
type LintMetadata struct {
	// Name is a lowercase underscore-separated string describing what a given
	// Lint checks. If Name beings with "w", the lint MUST NOT return Error, only
	// Warn. If Name beings with "e", the Lint MUST NOT return Warn, only Error.
	Name string `json:"name,omitempty"`

	// A human-readable description of what the Lint checks. Usually copied
	// directly from the CA/B Baseline Requirements or RFC 5280.
	Description string `json:"description,omitempty"`

	// The source of the check, e.g. "BRs: 6.1.6" or "RFC 5280: 4.1.2.6".
	Citation string `json:"citation,omitempty"`

	// Programmatic source of the check, BRs, RFC5280, or ZLint
	Source LintSource `json:"source"`

	// Lints automatically returns NE for all certificates where CheckApplies() is
	// true but with NotBefore < EffectiveDate. This check is bypassed if
	// EffectiveDate is zero. Please see CheckEffective for more information.
	EffectiveDate time.Time `json:"-"`

	// Lints automatically returns NE for all certificates where CheckApplies() is
	// true but with NotBefore >= IneffectiveDate. This check is bypassed if
	// IneffectiveDate is zero. Please see CheckEffective for more information.
	IneffectiveDate time.Time `json:"-"`
}

type CertificateLintInterface interface {
	CheckProfile(c *x509.Certificate) bool
	CheckApplies(c *x509.Certificate) bool
	Execute(c *x509.Certificate) *LintResult
}
