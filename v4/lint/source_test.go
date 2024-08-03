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

package lint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

// TestLintSourceMarshal tests that a LintSource can be correctly marshaled and
// unmarshalled.
func TestLintSourceMarshal(t *testing.T) {
	//nolint:musttag
	throwAway := struct {
		Source LintSource
	}{
		Source: Community,
	}

	jsonBytes, err := json.Marshal(&throwAway)
	if err != nil {
		t.Fatalf("failed to marshal LintSource: %v", err)
	}

	expectedJSON := fmt.Sprintf(`{"Source":%q}`, Community)
	if !bytes.Equal(jsonBytes, []byte(expectedJSON)) {
		t.Fatalf("expected JSON %q got %q", expectedJSON, string(jsonBytes))
	}

	err = json.Unmarshal(jsonBytes, &throwAway)
	if err != nil {
		t.Fatalf("err unmarshalling prev. marshaled LintSource: %v", err)
	}
	if throwAway.Source != Community {
		t.Fatalf("expected post-unmarshal value of %q got %q", Community, throwAway.Source)
	}

	badJSON := []byte(`{"Source":"cpu"}`)
	err = json.Unmarshal(badJSON, &throwAway)
	if err == nil {
		t.Fatalf("expected err unmarshalling bad LintSource value. Got nil")
	}
	if throwAway.Source != UnknownLintSource {
		t.Fatalf("expected Source to be %q after bad unmarshal, got %q\n", UnknownLintSource, throwAway.Source)
	}
}
