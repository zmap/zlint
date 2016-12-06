/* 	namesyntax.go
 	contains helper function for evaluating if a DNS name is in the preferred
	name syntax specified by RFC 1034 section 3.5
*/
package util

import (
	"regexp"
)

func IsInPrefSyn(name string) bool {
	// If the DNS name is just a space, it is valid
	if name == " " {
		return true
	}
	// This is the expression that matches the ABNF syntax from RFC 1034: Sec 3.5, specifically for subdomain since the " " case for domain is covered above
	prefsyn := regexp.MustCompile(`^([[:alpha:]]{1}(([[:alnum:]]|[-])*[[:alnum:]]{1})*){1}([.][[:alpha:]]{1}(([[:alnum:]]|[-])*[[:alnum:]]{1})*)*$`)
	return prefsyn.MatchString(name)
}
