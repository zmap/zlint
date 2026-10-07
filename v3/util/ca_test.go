package util

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

import (
	"testing"

	"github.com/zmap/zcrypto/encoding/asn1"
	"github.com/zmap/zcrypto/x509"
)

// TestIsServerAuthCertEidasNonWebQualified covers
// https://github.com/zmap/zlint/issues/951: an eIDAS Qualified Certificate
// with no EKU extension at all (RFC 5280's "valid for every purpose"
// default) but carrying a non-website QCP policy OID -- e.g. QCP-n-qscd, as
// used for qualified electronic signatures on a QSCD -- should not be
// treated as a TLS server certificate.
func TestIsServerAuthCertEidasNonWebQualified(t *testing.T) {
	testCases := []struct {
		name              string
		policyIdentifiers []asn1.ObjectIdentifier
		expected          bool
	}{
		{
			name:              "no EKU, no policies at all",
			policyIdentifiers: nil,
			expected:          true,
		},
		{
			name:              "no EKU, QCP-n-qscd only (issue 951 reproduction)",
			policyIdentifiers: []asn1.ObjectIdentifier{QCPnqscdPolicyOID},
			expected:          false,
		},
		{
			name:              "no EKU, QCP-n only",
			policyIdentifiers: []asn1.ObjectIdentifier{QCPnPolicyOID},
			expected:          false,
		},
		{
			name:              "no EKU, QCP-l-qscd only",
			policyIdentifiers: []asn1.ObjectIdentifier{QCPlqscdPolicyOID},
			expected:          false,
		},
		{
			name:              "no EKU, QCP-l only",
			policyIdentifiers: []asn1.ObjectIdentifier{QCPlPolicyOID},
			expected:          false,
		},
		{
			name:              "no EKU, QNCP-w-gen (a QWAC/website policy) only",
			policyIdentifiers: []asn1.ObjectIdentifier{QNCPwgenPolicyOID},
			expected:          true,
		},
		{
			name:              "no EKU, both QCP-n-qscd and QNCP-w-gen -- website policy wins",
			policyIdentifiers: []asn1.ObjectIdentifier{QCPnqscdPolicyOID, QNCPwgenPolicyOID},
			expected:          true,
		},
		{
			name:              "no EKU, unrelated policy OID",
			policyIdentifiers: []asn1.ObjectIdentifier{BRDomainValidatedOID},
			expected:          true,
		},
		{
			// A real-world regression case: some certificates in the wild
			// carry both a CABF BR reserved policy OID and an eIDAS
			// non-website QCP OID with no EKU extension. The BR policy OID
			// affirmatively puts the cert back in scope for the CABF BRs
			// (and thus server auth), overriding the eIDAS carve-out.
			name:              "no EKU, both a BR reserved policy and QCP-n-qscd -- BR policy wins",
			policyIdentifiers: []asn1.ObjectIdentifier{BRDomainValidatedOID, QCPnqscdPolicyOID},
			expected:          true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cert := &x509.Certificate{PolicyIdentifiers: tc.policyIdentifiers}
			if got := IsServerAuthCert(cert); got != tc.expected {
				t.Errorf("IsServerAuthCert() = %v, want %v", got, tc.expected)
			}
		})
	}
}

// TestIsServerAuthCertEidasNonWebQualifiedExplicitEKU confirms that an
// explicit serverAuth (or anyExtendedKeyUsage) EKU always wins, regardless of
// any eIDAS non-website QCP policy also present -- the carve-out only
// applies to the "no EKU extension at all" default.
func TestIsServerAuthCertEidasNonWebQualifiedExplicitEKU(t *testing.T) {
	cert := &x509.Certificate{
		ExtKeyUsage:       []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		PolicyIdentifiers: []asn1.ObjectIdentifier{QCPnqscdPolicyOID},
	}
	if !IsServerAuthCert(cert) {
		t.Errorf("IsServerAuthCert() = false, want true: an explicit serverAuth EKU must not be overridden by a QC policy OID")
	}
}
