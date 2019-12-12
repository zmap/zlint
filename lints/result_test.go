package lints

import (
	"encoding/json"
	"testing"
)

func TestMarshalingLintStatus(t *testing.T) {
	testCases := []struct {
		result       LintStatus
		expectedJSON string
	}{
		{
			result:       Reserved,
			expectedJSON: `"reserved"`,
		},
		{
			result:       NA,
			expectedJSON: `"NA"`,
		},
		{
			result:       NE,
			expectedJSON: `"NE"`,
		},
		{
			result:       Pass,
			expectedJSON: `"pass"`,
		},
		{
			result:       Notice,
			expectedJSON: `"info"`,
		},
		{
			result:       Warn,
			expectedJSON: `"warn"`,
		},
		{
			result:       Error,
			expectedJSON: `"error"`,
		},
		{
			result:       Fatal,
			expectedJSON: `"fatal"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.result.String(), func(t *testing.T) {
			j, err := json.Marshal(tc.result)
			if err != nil {
				t.Error("Failed to marshal LintStatus")
			}
			if string(j) != tc.expectedJSON {
				t.Errorf("Expected LintStatus to marshal to JSON %q, got %q",
					tc.expectedJSON,
					j)
			}
			var in LintStatus
			if err := json.Unmarshal(j, &in); err != nil {
				t.Errorf("Expected to unmarshal %q without error. Got %v", j, err)
			}
			if in != tc.result {
				t.Errorf("Expected to unmarshal %q to %#v, got %#v", j, tc.result, in)
			}
		})
	}

}
