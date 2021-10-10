/*
 * ZLint Copyright 2021 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package util

import (
	"regexp"

	"golang.org/x/net/idna"
)

var acePrefix = regexp.MustCompile(`^[xX][nN]--.*`)

// IsACEPrefixed checks whether-or-not the given string is prefixed
// with the case-insensitive string "xn--".
//
// This check is useful given the bug following bug report for IDNA wherein
// the ACE prefix incorrectly taken to be case-sensitive.
//
// https://github.com/golang/go/issues/48778
func IsACEPrefixed(s string) bool {
	return acePrefix.MatchString(s)
}

// IdnaToUnicode is a wrapper around idna.ToUnicode.
//
// If the provided string is ACE prefixed, then that ACE prefix
// is coerced to a lowercase "xn--" before processing by
// then idna package.
//
// This is only necessary due to the bug at https://github.com/golang/go/issues/48778
func IdnaToUnicode(s string) (string, error) {
	if IsACEPrefixed(s) {
		s = "xn--" + s[4:]
	}
	return idna.ToUnicode(s)
}
