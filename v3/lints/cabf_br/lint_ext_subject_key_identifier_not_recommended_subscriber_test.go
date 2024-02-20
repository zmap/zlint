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

func TestSubjectKeyIdNotRecommendedSubscriber(t *testing.T) {
	type Test struct {
		input string
		want  lint.LintStatus
	}
	data := []Test{
		{
			input: "warn_subject_key_identifier_not_recommended_subscriber.pem",
			want:  lint.Warn,
		},
		{
			input: "pass_subject_key_identifier_not_recommended_subscriber.pem",
			want:  lint.Pass,
		},
		{
			input: "ne_subject_key_identifier_not_recommended_subscriber.pem",
			want:  lint.NE,
		},
	}
	for _, in := range data {
		in := in
		t.Run(in.input, func(t *testing.T) {
			out := test.TestLint("w_ext_subject_key_identifier_not_recommended_subscriber", in.input)
			if out.Status != in.want {
				t.Errorf("expected %s, got %s", in.want, out.Status)
			}
		})
	}
}
