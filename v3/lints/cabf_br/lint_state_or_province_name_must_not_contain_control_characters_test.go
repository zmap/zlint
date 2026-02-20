/*
 * ZLint Copyright 2024 Regents of the University of Michigan
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

package cabf_br

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

func TestStateOrProvinceNameMustNotContainControlCharacters(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected lint.LintStatus
	}{
		{
			name:     "Valid state name",
			input:    "stateWithValidName.pem",
			expected: lint.Pass,
		},
		{
			name:     "State with equals",
			input:    "stateWithEquals.pem",
			expected: lint.Error,
		},
		{
			name:     "State with colon",
			input:    "stateWithColon.pem",
			expected: lint.Error,
		},
		{
			name:     "State with open brace",
			input:    "stateWithOpenBrace.pem",
			expected: lint.Error,
		},
		{
			name:     "State with close brace",
			input:    "stateWithCloseBrace.pem",
			expected: lint.Error,
		},
		{
			name:     "State with open bracket",
			input:    "stateWithOpenBracket.pem",
			expected: lint.Error,
		},
		{
			name:     "State with close bracket",
			input:    "stateWithCloseBracket.pem",
			expected: lint.Error,
		},
		{
			name:     "State with semicolon",
			input:    "stateWithSemicolon.pem",
			expected: lint.Error,
		},
		{
			name:     "State with quote",
			input:    "stateWithQuote.pem",
			expected: lint.Error,
		},
		{
			name:     "State with pipe",
			input:    "stateWithPipe.pem",
			expected: lint.Error,
		},
		{
			name:     "State with backslash",
			input:    "stateWithBackslash.pem",
			expected: lint.Error,
		},
		{
			name:     "Valid locality name",
			input:    "localityWithValidName.pem",
			expected: lint.Pass,
		},
		{
			name:     "Locality with equals",
			input:    "localityWithEquals.pem",
			expected: lint.Error,
		},
		{
			name:     "Locality with colon",
			input:    "localityWithColon.pem",
			expected: lint.Error,
		},
		{
			name:     "Locality with open brace",
			input:    "localityWithOpenBrace.pem",
			expected: lint.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := test.TestLint("e_state_or_province_name_must_not_contain_control_characters", tt.input)
			if out.Status != tt.expected {
				t.Errorf("%s: expected %s, got %s", tt.input, tt.expected, out.Status)
			}
		})
	}
}
