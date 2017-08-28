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
	rightLabel = strings.ToUpper(rightLabel)
	if _, ok := tldMap[rightLabel]; ok {
		return true
	} else {
		return false
	}
}
