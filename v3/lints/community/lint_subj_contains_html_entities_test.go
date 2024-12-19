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

package community

import (
	"testing"

	"github.com/zmap/zlint/v3/lint"
	"github.com/zmap/zlint/v3/test"
)

/*
   TEST CASES

   File                Result      Description
   ===============     ======      ===========
   html_entity_ok1     Pass        Clean certificate (no HTML entities)
   html_entity_ok2     Pass        With a pattern that resembles, but is not, an HTML entity
   html_entity_ok3     Pass        With an HTML entity, but lint is bypassed via configuration
   html_entity_ko1     Error       HTML entity in organization
   html_entity_ko2     Error       HTML entity in stateOrProvince (Turks & Caicos Islands)
   html_entity_ko3     Error       HTML entity in locality (La Roque-d'Anthéron)
*/

func TestSubjectContainsHTMLEntities(t *testing.T) {

	type Data struct {
		input  string
		config string
		want   lint.LintStatus
	}

	data := []Data{
		{
			input: "html_entity_ok1.pem",
			want:  lint.Pass,
		},
		{
			input: "html_entity_ok2.pem",
			want:  lint.Pass,
		},
		{
			input: "html_entity_ok3.pem",
			config: `
                [e_subj_contains_html_entities]
                Skip = true
                `,
			want: lint.Pass,
		},
		{
			input: "html_entity_ko1.pem",
			want:  lint.Error,
		},
		{
			input: "html_entity_ko2.pem",
			want:  lint.Error,
		},
		{
			input: "html_entity_ko3.pem",
			want:  lint.Error,
		},
	}
	for _, testData := range data {
		testData := testData
		t.Run(testData.input, func(t *testing.T) {
			out := test.TestLintWithConfig("e_subj_contains_html_entities", testData.input, testData.config)
			if out.Status != testData.want {
				t.Errorf("expected %s, got %s", testData.want, out.Status)
			}
		})
	}

}
