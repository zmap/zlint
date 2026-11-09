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
	"testing"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zcrypto/x509/pkix"
	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
	"github.com/zmap/zlint/v3/util"
)

/*
 * Explanation of test file names:
 *
 *      nup(1|0) = CRL nextUpdate is present (1) or absent (0)
 *      sub(1|0) = CRL covers (1) Subscriber certificates or not (0)
 *      len(1|0) = CRL lifespan is within (0) or beyond (1) the limit set by BRs
 *      eff(1|0) = CRL thisUpdate is before (0) this lint's effective date or after it (1)
 */

func TestCrlNextUpdateInvalid(t *testing.T) {

	type Data struct {
		input  string
		config string
		want   lint.LintStatus
	}
	data := []Data{
		{
			input: "crl_nextupdate_nup1_sub1_len0_eff0.pem",
			want:  lint.Pass,
		},
		{
			input: "crl_nextupdate_nup1_sub1_len1_eff0.pem",
			want:  lint.NE,
		},
		{
			input: "crl_nextupdate_nup1_sub1_len1_eff1.pem",
			want:  lint.Error,
		},
		{
			input: "crl_nextupdate_nup1_sub0_len0_eff0.pem",
			want:  lint.Pass,
			config: `
[e_crl_next_update_invalid]
SubscriberCRL = false`,
		},
		{
			input: "crl_nextupdate_nup1_sub0_len1_eff0.pem",
			want:  lint.NE,
			config: `
[e_crl_next_update_invalid]
SubscriberCRL = false`,
		},
		{
			input: "crl_nextupdate_nup1_sub0_len1_eff1.pem",
			want:  lint.Error,
			config: `
[e_crl_next_update_invalid]
SubscriberCRL = false`,
		},
		{
			input: "crl_nextupdate_nup0_sub0_len0_eff0.pem",
			want:  lint.NA,
		},
		// The following cases carry an Issuing Distribution Point that scopes
		// the CRL to CA or subscriber certificates. That in-band signal must
		// take precedence over the SubscriberCRL configuration default.
		{
			// IDP: onlyContainsCACerts; lifespan beyond 10 days but within
			// 12 months -> Pass without any configuration.
			input: "crl_nextupdate_idp_ca_eff1.pem",
			want:  lint.Pass,
		},
		{
			// IDP: onlyContainsCACerts; lifespan beyond 12 months -> still an
			// Error, proving the CA limit is really enforced.
			input: "crl_nextupdate_idp_ca_beyond12m_eff1.pem",
			want:  lint.Error,
		},
		{
			// IDP: onlyContainsUserCerts; lifespan beyond 10 days -> Error.
			input: "crl_nextupdate_idp_ee_eff1.pem",
			want:  lint.Error,
		},
		{
			// Config says CA, but the IDP says subscriber certs; the IDP wins.
			input: "crl_nextupdate_idp_ee_eff1.pem",
			config: `
[e_crl_next_update_invalid]
SubscriberCRL = false`,
			want: lint.Error,
		},
	}

	for _, testData := range data {
		testData := testData
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestRevocationListLintWithConfig(t, "e_crl_next_update_invalid", testData.input, testData.config)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}
}

// mustMarshalIDP encodes an Issuing Distribution Point carrying only the scope
// booleans given, so that idpCRLScope can be exercised without a signed CRL.
func mustMarshalIDP(t *testing.T, idp interface{}) []byte {
	t.Helper()
	der, err := asn1.Marshal(idp)
	if err != nil {
		t.Fatalf("marshaling IDP: %v", err)
	}
	return der
}

func crlWithIDP(idp []byte) *x509.RevocationList {
	return &x509.RevocationList{
		Extensions: []pkix.Extension{{
			Id:       util.IssuingDistOID,
			Critical: true,
			Value:    idp,
		}},
	}
}

// TestIDPCRLScope covers the scoping signal itself, including the cases that
// cannot be expressed with a signed CRL fixture: an IDP asserting nothing, an
// IDP asserting both scopes (which BR §7.2.2.1 prohibits), and an unparseable
// IDP.
func TestIDPCRLScope(t *testing.T) {
	caOnly := mustMarshalIDP(t, struct {
		OnlyContainsCACerts bool `asn1:"optional,tag:2"`
	}{OnlyContainsCACerts: true})

	userOnly := mustMarshalIDP(t, struct {
		OnlyContainsUserCerts bool `asn1:"optional,tag:1"`
	}{OnlyContainsUserCerts: true})

	bothScopes := mustMarshalIDP(t, struct {
		OnlyContainsUserCerts bool `asn1:"optional,tag:1"`
		OnlyContainsCACerts   bool `asn1:"optional,tag:2"`
	}{OnlyContainsUserCerts: true, OnlyContainsCACerts: true})

	data := []struct {
		name      string
		idp       []byte
		wantScope crlScope
		wantErr   bool
		wantErrIs error
	}{
		{
			name:      "onlyContainsCACerts",
			idp:       caOnly,
			wantScope: crlScopeCA,
		},
		{
			name:      "onlyContainsUserCerts",
			idp:       userOnly,
			wantScope: crlScopeSubscriber,
		},
		{
			// An IDP may be present for other reasons (e.g. a distribution
			// point) without scoping the CRL to either kind of certificate.
			name:      "no scope asserted",
			idp:       []byte{0x30, 0x00},
			wantScope: crlScopeUnknown,
		},
		{
			name:      "both scopes asserted",
			idp:       bothScopes,
			wantScope: crlScopeUnknown,
			wantErr:   true,
			wantErrIs: errAmbiguousIDPScope,
		},
		{
			name:      "unparseable IDP",
			idp:       []byte{0x02, 0x01, 0x01},
			wantScope: crlScopeUnknown,
			wantErr:   true,
		},
	}

	for _, testData := range data {
		testData := testData
		t.Run(testData.name, func(t *testing.T) {
			scope, err := idpCRLScope(crlWithIDP(testData.idp))
			if scope != testData.wantScope {
				t.Errorf("expected scope %d, got %d", testData.wantScope, scope)
			}
			if (err != nil) != testData.wantErr {
				t.Fatalf("expected error %t, got %v", testData.wantErr, err)
			}
			if testData.wantErrIs != nil && !errors.Is(err, testData.wantErrIs) {
				t.Errorf("expected error %v, got %v", testData.wantErrIs, err)
			}
		})
	}
}

// TestIDPCRLScopeAbsent covers the fallback path: no IDP extension at all.
func TestIDPCRLScopeAbsent(t *testing.T) {
	scope, err := idpCRLScope(&x509.RevocationList{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if scope != crlScopeUnknown {
		t.Errorf("expected scope %d, got %d", crlScopeUnknown, scope)
	}
}
