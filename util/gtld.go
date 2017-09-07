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

func IsInTLDMap(label string) bool {
	label = strings.ToUpper(label)
	if _, ok := tldMap[label]; ok {
		return true
	} else {
		return false
	}
}
