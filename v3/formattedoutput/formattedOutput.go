package formattedoutput

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/zmap/zlint/v3"
	"github.com/zmap/zlint/v3/lint"
)

type resultsTable struct {
	resultCount              map[lint.LintStatus]int
	resultDetails            map[lint.LintStatus][]string
	lintLevelsAboveThreshold map[int]lint.LintStatus
	sortedLevels             []int
}

func (r *resultsTable) newRT(threshold lint.LintStatus, results *zlint.ResultSet, longSummary bool) resultsTable {

	r.resultCount = make(map[lint.LintStatus]int)
	r.resultDetails = make(map[lint.LintStatus][]string)
	r.lintLevelsAboveThreshold = make(map[int]lint.LintStatus)

	// Make the list of lint levels that matter
	for _, i := range lint.StatusLabelToLintStatus {
		if i <= threshold {
			continue
		}
		r.lintLevelsAboveThreshold[int(i)] = i
	}
	// Set all of the levels to 0 events so they are all displayed
	// in the -summary table
	for _, level := range r.lintLevelsAboveThreshold {
		r.resultCount[level] = 0
	}
	// Count up the number of each event
	for lintName, lintResult := range results.Results {
		if lintResult.Status > threshold {
			r.resultCount[lintResult.Status]++
			if longSummary {
				r.resultDetails[lintResult.Status] = append(
					r.resultDetails[lintResult.Status],
					lintName,
				)
			}
		}
	}
	// Sort the levels we have so we can get a nice output
	for key := range r.resultCount {
		r.sortedLevels = append(r.sortedLevels, int(key))
	}
	sort.Ints(r.sortedLevels)

	return *r
}

func OutputSummary(zlintResult *zlint.ResultSet, longSummary bool) {
	// Set the threashold under which (inclusive) events are not
	// counted
	threshold := lint.Pass

	rt := (&resultsTable{}).newRT(threshold, zlintResult, longSummary)

	// make and print the requested table type
	if longSummary {
		// make a table with the internal lint names grouped
		// by type
		var olsl string
		headings := []string{
			"Level",
			"# occurrences",
			"                      Details                      ",
		}
		lines := [][]string{}
		lsl := ""
		rescount := ""

		hlengths := printTableHeadings(headings)
		// Construct the table lines, but don't repeat
		// LintStatus(level) or the results count.  Also, just
		// because a level wasn't seen doesn't mean it isn't
		// important; display "empty" levels, too
		for _, level := range rt.sortedLevels {
			foundDetail := false
			for _, detail := range rt.resultDetails[lint.LintStatus(level)] {
				if lint.LintStatus(level).String() != olsl {
					olsl = lint.LintStatus(level).String()
					lsl = olsl
					rescount = strconv.Itoa(rt.resultCount[lint.LintStatus(level)])
				} else {
					lsl = ""
					rescount = ""
				}
				lines = append(lines, ([]string{lsl, rescount, detail}))
				foundDetail = true
			}
			if !foundDetail {
				lines = append(lines, []string{
					lint.LintStatus(level).String(),
					strconv.Itoa(rt.resultCount[lint.LintStatus(level)]),
					" - ",
				})
			}
		}
		printTableBody(hlengths, lines)
	} else {
		headings := []string{"Level", "# occurrences"}
		hlengths := printTableHeadings(headings)
		lines := [][]string{}
		for _, level := range rt.sortedLevels {
			lines = append(lines, []string{
				lint.LintStatus(level).String(),
				strconv.Itoa(rt.resultCount[lint.LintStatus(level)])})
		}
		printTableBody(hlengths, lines)
		fmt.Printf("\n")
	}
}

func printTableHeadings(headings []string) []int {
	hlengths := []int{}
	for i, h := range headings {
		hlengths = append(
			hlengths,
			utf8.RuneCountInString(h)+1)
		fmt.Printf("| %s ", strings.ToUpper(h))
		if i == len(headings)-1 {
			fmt.Printf("|\n")
			for ii, j := range hlengths {
				fmt.Printf("+%s", strings.Repeat("-", j+1))
				if ii == len(headings)-1 {
					fmt.Printf("+\n")
				}
			}
		}
	}
	return hlengths
}

func printTableBody(hlengths []int, lines [][]string) {
	for _, line := range lines {
		for i, hlen := range hlengths {
			// This makes a format string with the
			// right widths, e.g. "%7.7s"
			fmtstring := fmt.Sprintf("|%%%[1]d.%[1]ds", hlen)
			fmt.Printf(fmtstring, line[i])
			if i == len(hlengths)-1 {
				fmt.Printf(" |\n")
			} else {
				fmt.Printf(" ")
			}
		}
	}

}
