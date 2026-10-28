
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

package util

import (
	"net"
	"testing"
)

// TestIsIANAReservedIPv6Documentation verifies that IsIANAReserved
// recognizes the IPv6 documentation prefix 3fff::/20 defined in
// RFC 9637, while preserving the existing RFC 3849 classification.
func TestIsIANAReservedIPv6Documentation(t *testing.T) {
	tests := []struct {
		name     string
		address  string
		expected bool
	}{
		{
			name:     "RFC 9637 first address",
			address:  "3fff::",
			expected: true,
		},
		{
			name:     "RFC 9637 address within prefix",
			address:  "3fff:123:4567::1",
			expected: true,
		},
		{
			name:     "RFC 9637 last address",
			address:  "3fff:fff:ffff:ffff:ffff:ffff:ffff:ffff",
			expected: true,
		},
		{
			name:     "Address immediately outside prefix",
			address:  "3fff:1000::",
			expected: false,
		},
		{
			name:     "Existing RFC 3849 documentation prefix",
			address:  "2001:db8::1",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip := net.ParseIP(tt.address)
			if ip == nil {
				t.Fatalf("Failed to parse IP address %s", tt.address)
			}

			actual := IsIANAReserved(ip)
			if actual != tt.expected {
				t.Errorf(
					"For IP address %s, expected %t, got %t",
					tt.address, tt.expected, actual,
				)
			}
		})
	}
}

// TestIntersectsIANAReservedIPv6Documentation verifies that
// IntersectsIANAReserved detects networks that intersect the
// RFC 9637 documentation prefix 3fff::/20.
func TestIntersectsIANAReservedIPv6Documentation(t *testing.T) {
	tests := []struct {
		name     string
		cidr     string
		expected bool
	}{
		{
			name:     "Exact RFC 9637 documentation prefix",
			cidr:     "3fff::/20",
			expected: true,
		},
		{
			name:     "Subnet within documentation prefix",
			cidr:     "3fff:100::/32",
			expected: true,
		},
		{
			name:     "Larger network containing documentation prefix",
			cidr:     "3fff::/19",
			expected: true,
		},
		{
			name:     "Network outside documentation prefix",
			cidr:     "3fff:1000::/20",
			expected: false,
		},
		{
			name:     "Existing RFC 3849 documentation prefix",
			cidr:     "2001:db8::/32",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, network, err := net.ParseCIDR(tt.cidr)
			if err != nil {
				t.Fatalf("Failed to parse CIDR %s: %v", tt.cidr, err)
			}

			actual := IntersectsIANAReserved(*network)
			if actual != tt.expected {
				t.Errorf(
					"For CIDR %s, expected %t, got %t",
					tt.cidr, tt.expected, actual,
				)
			}
		})
	}
}
