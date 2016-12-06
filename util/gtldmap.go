/* dataupdate.go
 * File used to parse newgtlds.csv and generate a map
 */

package util

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
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
	if theMap[theGTLD] {
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

func parseData() ([]string, int) {

	// open a file
	file, err := os.Open("../data/newgtlds.txt")
	if err != nil {
		fmt.Println(err)
		return nil, 0
	}

	// make sure it gets closed
	defer file.Close()

	// create a new scanner and read the file line by line
	scanner := bufio.NewScanner(file)

	var theTlds []string

	for scanner.Scan() {
		theTlds = append(theTlds, scanner.Text())
	}

	return theTlds[1:], len(theTlds) //Shave off the header

	/*
		    // Load the csv from it's relative location
		    f, _ := os.Open("data/newgtlds.csv")

		    // Create a new reader.
		    r := csv.NewReader(bufio.NewReader(f))

		    // I'm certain there is a better way to do this, but I kept getting weird errors
		    r.Read() // Header line
		    r.Read() // Column definition line
		    r.FieldsPerRecord = 6 // It defaulted to 1 because of the header line, correct
		    r.LazyQuotes = true // Not sure why this is off by default

		    // Read all of the remaining data into a [][]string
			records, err := r.ReadAll()
			if err != nil {
				fmt.Println(err)
			}

		    var theTlds []string

		    // Grab only the first element (the tld) from each row
			for _, row := range records {
		        theTlds = append(theTlds, row[0])
		    }
		    return theTlds, len(theTlds)
	*/
}
