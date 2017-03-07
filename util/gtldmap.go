/* dataupdate.go
 * File used to parse newgtlds.csv and generate a map
 */

package util

import (
	"net"
	"regexp"
	"strings"
	"net/http"
	"io/ioutil"
)

var tldMap map[string]bool

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

	theMap := fetchTLDMap()

	if _, ok := theMap[theGTLD]; ok {
		return 1
	} else {
		return 0
	}

}

func fetchTLDMap() map[string]bool {
	if tldMap != nil {
		return tldMap
	}
	entries, datSize := parseData()
	tldMap = make(map[string]bool, datSize)

	for _, entry := range entries {
		tldMap[strings.ToUpper(entry)] = true
	}

	return tldMap

}

func parseData()  ([]string, int){
	//Read data from IANA
	url := "http://data.iana.org/TLD/tlds-alpha-by-domain.txt"
	resp, err := http.Get(url)
	if err != nil {
		//handle error
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var tld_data []string
	var temp string = ""

	for i := 0; i < len(body); i++{
		if body[i] == '\n'{
			tld_data = append(tld_data, temp)
			temp = ""
			continue
		}
		temp += string(body[i])
	}
	return tld_data[1:], len(tld_data)
}
