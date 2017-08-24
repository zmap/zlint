package util

import (
	"net"
	"regexp"
	"strings"
)

func removeQuestionMarks(domain string) string {
	hasQuestionMarkPrefix := strings.HasPrefix(domain, "?.")
	for hasQuestionMarkPrefix == true {
		domain = domain[2:]
		hasQuestionMarkPrefix = strings.HasPrefix(domain, "?.")
	}
	return domain
}

func checkFQDN(domain string) bool {
	dnsLabel := "([A-Za-z0-9*_-]{1,63})"
	fqdn := `\A(` + dnsLabel + `\.)*` + dnsLabel + `\z`
	re := regexp.MustCompile(fqdn)
	matched := re.MatchString(domain)
	return matched
}

func IsFQDN(domain string) bool {
	domain = removeQuestionMarks(domain)
	if strings.HasPrefix(domain, "*.") {
		domain = domain[2:]
	}
	return checkFQDN(domain)
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
		return ""
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
	if IsFQDN(auth) {
		return true
	}
	if net.ParseIP(auth) != nil {
		return true
	}
	return false
}
