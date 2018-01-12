/*
 * ZLint Copyright 2018 Regents of the University of Michigan
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
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/weppos/publicsuffix-go/publicsuffix"
	"github.com/zmap/zcrypto/x509"
)

func RemovePrependedQuestionMarks(domain string) string {
	for strings.HasPrefix(domain, "?.") {
		domain = domain[2:]
	}
	return domain
}

func RemovePrependedWildcard(domain string) string {
	if strings.HasPrefix(domain, "*.") {
		domain = domain[2:]
	}
	return domain
}

func IsFQDN(domain string) bool {
	domain = RemovePrependedWildcard(domain)
	domain = RemovePrependedQuestionMarks(domain)
	return govalidator.IsURL(domain)
}

func GetAuthority(uri string) string {
	if len(uri) < 4 {
		return ""
	}
	idx := strings.Index(uri, "//")
	for i := idx + 2; i < len(uri); i++ {
		if uri[i] == '/' || uri[i] == '#' || uri[i] == '?' {
			return uri[idx+2 : i]
		}
	}
	if idx != -1 {
		return uri[idx+2:]
	}
	return ""
}

func GetHost(auth string) string {
	begin := strings.Index(auth, "@")
	if begin == -1 || begin == len(auth)-1 {
		begin = -1
	}
	end := strings.Index(auth, ":")
	if end == -1 {
		end = len(auth)
	}
	if end < begin {
		return ""
	}
	return auth[begin+1 : end]
}

func AuthIsFQDNOrIP(auth string) bool {
	return IsFQDNOrIP(GetHost(auth))
}

func IsFQDNOrIP(host string) bool {
	if IsFQDN(host) {
		return true
	}
	if net.ParseIP(host) != nil {
		return true
	}
	return false
}

func DNSNamesExist(cert *x509.Certificate) bool {
	if cert.Subject.CommonName == "" && len(cert.DNSNames) == 0 {
		return false
	} else {
		return true
	}
}

func ICANNPublicSuffixParse(domain string) (*publicsuffix.DomainName, error) {
	return publicsuffix.ParseFromListWithOptions(publicsuffix.DefaultList, domain, &publicsuffix.FindOptions{IgnorePrivate: true, DefaultRule: publicsuffix.DefaultRule})
}

func CommonNameIsIP(cert *x509.Certificate) bool {
	ip := net.ParseIP(cert.Subject.CommonName)
	if ip == nil {
		return false
	} else {
		return true
	}
}
