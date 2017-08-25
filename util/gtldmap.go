/* dataupdate.go
 * File used to parse newgtlds.csv and generate a map
 */

package util

import (
	"net"
	"regexp"
	"strings"
)

/* Return schema for IsValidTLD
 * 1 = valid TLD
 * 0 = TLD not in list
 * -1 = input string not recognized as DNS
 */

func IsValidGTLD(input string) int {
	//First make sure it isn't an IP, which gets return: -1
	if input == "" || net.ParseIP(input) != nil {
		return -1
	}
	//Finding if given string contains a gtld
	containsgtld := regexp.MustCompile(`[.]([-]|[[:alnum:]]){2,63}$`)
	isHost := containsgtld.MatchString(input)

	if !isHost {
		return -1
	}

	theGTLD := strings.ToUpper(containsgtld.FindString(input)[1:]) //Shave off the .

	if _, ok := tldMap[theGTLD]; ok {
		return 1
	} else {
		return 0
	}

}
