/*
 * ZLint Copyright 2023 Regents of the University of Michigan
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

	"golang.org/x/net/idna"
)

func TestHasReservedLabelPrefix(t *testing.T) {
	input := map[string]bool{
		"ab--":       true,
		"ab--foo":    true,
		"a---foo":    true,
		"A---foo":    true,
		"XN--foo":    true,
		"":           false,
		"a-b":        false,
		"a--":        false,
		"foobar--aa": false,
		"XNA--foo":   false,
	}
	for input, want := range input {
		got := HasReservedLabelPrefix(input)
		if got != want {
			t.Errorf("got %v want %v for input '%s'", got, want, input)
		}
	}
}

func TestHasXNLabelPrefix(t *testing.T) {
	input := map[string]bool{
		"xn--zlint.org": true,
		"Xn--zlint.org": true,
		"xN--zlint.org": true,
		"XN--zlint.org": true,
		"xn--":          true,
		"Xn--":          true,
		"xN--":          true,
		"XN--":          true,
		"-xn--":         false,
		"-Xn--":         false,
		"-xN--":         false,
		"-XN--":         false,
		"":              false,
	}
	for input, want := range input {
		got := HasXNLabelPrefix(input)
		if got != want {
			t.Errorf("got %v want %v for input '%s'", got, want, input)
		}
	}
}

func TestIdnaToUnicode(t *testing.T) {
	type testData struct {
		input   string
		want    string
		wantErr bool
	}
	input := []testData{
		{"xn--Mnchen-Ost-9db", "München-Ost", false},
		{"xn--Mnchen-ost-9db", "München-ost", false},
		{"xn--", "", false},
		{"xN--12311613412431243.com", "xn--12311613412431243.com", true},

		{"Xn--Mnchen-Ost-9db", "München-Ost", false},
		{"Xn--Mnchen-ost-9db", "München-ost", false},
		{"Xn--", "", false},
		{"xN--12311613412431243.com", "xn--12311613412431243.com", true},

		{"xN--Mnchen-Ost-9db", "München-Ost", false},
		{"xN--Mnchen-ost-9db", "München-ost", false},
		{"xN--", "", false},
		{"xN--12311613412431243.com", "xn--12311613412431243.com", true},

		{"XN--Mnchen-Ost-9db", "München-Ost", false},
		{"XN--Mnchen-ost-9db", "München-ost", false},
		{"XN--", "", false},
		{"xN--12311613412431243.com", "xn--12311613412431243.com", true},
	}
	for _, data := range input {
		got, err := IdnaToUnicode(data.input)
		gotErr := err != nil
		if gotErr != data.wantErr || data.want != got {
			t.Errorf("got string '%s' error '%v' for test data %v", got, err, data)
		}
	}
}

// This test checks whether or not https://github.com/golang/go/issues/48778 ever got fixed.
// If it did then we can likely delete some code from utils since we don't have to handle it
// with kid gloves anymore.
func TestIdnaToUnicodeBugIsStillThere(t *testing.T) {
	s, _ := idna.ToUnicode("Xn--Mnchen-Ost-9db")
	if s == "München-Ost" {
		t.Fatal("https://github.com/golang/go/issues/48778 appears to have been fixed")
	}
}
