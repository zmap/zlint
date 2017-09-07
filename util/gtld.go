/* dataupdate.go
 * File used to parse newgtlds.csv and generate a map
 */

package util

import (
	"strings"
)

func HasValidTLD(domain string) bool {
	labels := strings.Split(domain, ".")
	rightLabel := labels[len(labels)-1]
	return IsInTLDMap(rightLabel)
}

func IsInTLDMap(fqdn string) bool {
	fqdn = strings.ToUpper(fqdn)
	if _, ok := tldMap[fqdn]; ok {
		return true
	} else {
		return false
	}
}
