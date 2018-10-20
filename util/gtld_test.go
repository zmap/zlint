/*
 * ZLint Copyright 2018 Regents of the University of Michigan
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
	"testing"
	"time"
)

func TestHasValidTLD(t *testing.T) {
	domain := "google.com"
	expected := true
	actual := HasValidTLD(domain, time.Now())
	if expected != actual {
		t.Error(
			"For", domain,
			"expected", expected,
			"got", actual,
		)
	}
}

func TestHasValidTLDUppercaseName(t *testing.T) {
	domain := "GOOGLE.COM"
	expected := true
	actual := HasValidTLD(domain, time.Now())
	if expected != actual {
		t.Error(
			"For", domain,
			"expected", expected,
			"got", actual,
		)
	}
}
