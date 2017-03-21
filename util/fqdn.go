package util

import (
	"net"
	"regexp"
	"strings"
)

func IsFQDN(domain string) bool {
	RegExp := regexp.MustCompile(`^((([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9])|\*)\.){2,}([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})(\.)*$`)
	return RegExp.MatchString(domain)
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
