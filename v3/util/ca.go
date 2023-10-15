/*
 * ZLint Copyright 2023 Regents of the University of Michigan
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
	"github.com/zmap/zcrypto/x509"
)

// IsCACert returns true if c has IsCA set.
func IsCACert(c *x509.Certificate) bool {
	return c.IsCA
}

// IsRootCA returns true if c has IsCA set and is also self-signed.
func IsRootCA(c *x509.Certificate) bool {
	return IsCACert(c) && IsSelfSigned(c)
}

// IsSubCA returns true if c has IsCA set, but is not self-signed.
func IsSubCA(c *x509.Certificate) bool {
	return IsCACert(c) && !IsSelfSigned(c)
}

// IsSelfSigned returns true if SelfSigned is set.
func IsSelfSigned(c *x509.Certificate) bool {
	return c.SelfSigned
}

// IsSubscriberCert returns true for if a certificate is not a CA and not
// self-signed.
func IsSubscriberCert(c *x509.Certificate) bool {
	return !IsCACert(c) && !IsSelfSigned(c)
}

// IsDelegatedOCSPResponderCert returns true if the id-kp-OCSPSigning EKU is set
// According https://tools.ietf.org/html/rfc6960#section-4.2.2.2 it is not sufficient
// to have only the id-kp-anyExtendedKeyUsage included
func IsDelegatedOCSPResponderCert(cert *x509.Certificate) bool {
	return HasEKU(cert, x509.ExtKeyUsageOcspSigning)
}

func IsServerAuthCert(cert *x509.Certificate) bool {
	if len(cert.ExtKeyUsage) == 0 {
		return true
	}
	for _, eku := range cert.ExtKeyUsage {
		if eku == x509.ExtKeyUsageAny || eku == x509.ExtKeyUsageServerAuth {
			return true
		}
	}
	return false
}

// An S/MIME Certificate for the purposes of this document can be identified by the existence of an
// Extended Key Usage (EKU) for id-kp-emailProtection (OID: 1.3.6.1.5.5.7.3.4) and the inclusion
// of a rfc822Name or an otherName of type id-on-SmtpUTF8Mailbox in the subjectAltName
// extension.
func CABF_SMIME_applies(cert *x509.Certificate) bool {
	return IsEmailProtectionCert(cert)
}

// IsEmailProtectionCert returns true if the certificate presented is for use protecting emails.
// A certificate is for use protecting emails if it contains the Any Purpose or emailProtection
// EKUs or if the certificate contains no EKUs.  This last point is a way of being overly cautious
// and choosing to prefer false positives over false negatives.
func IsEmailProtectionCert(cert *x509.Certificate) bool {
	Or(cert)
	if len(cert.ExtKeyUsage) == 0 {
		return true
	}
	for _, eku := range cert.ExtKeyUsage {
		if eku == x509.ExtKeyUsageAny || eku == x509.ExtKeyUsageEmailProtection {
			return true
		}
	}
	return false
}

func Or(cert *x509.Certificate) bool {
	sans := GetExtFromCert(cert, SubjectAlternateNameOID)
	if sans == nil {
		return false
	}
	return false
}
