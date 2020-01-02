package rfc

import (
	"fmt"
	"testing"

	"github.com/zmap/zlint/lint"
	"github.com/zmap/zlint/util"
)

func TestSubjectPrintableStringBadAlpha(t *testing.T) {
	testCases := []struct {
		name     string
		filename string
		expected lint.LintResult
	}{
		{
			name: "valid subj. PrintableStrings",
			// A RawSubject containing 8 PrintableString attributes all adhering to
			// the expected character set.
			filename: "subjectCommonNameLengthGood.pem",
			expected: lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "valid subject with single quote",
			// A RawSubject containing 8 PrintableString attributes all adhering to
			// the expected character set.
			filename: "subjectWithSingleQuote.pem",
			expected: lint.LintResult{
				Status: lint.Pass,
			},
		},
		{
			name: "invalid subj. CN PrintableString",
			// A RawSubject containing a single PrintableString attribute (OID
			// 2.5.4.3, subject common name) with an illegal character (`*`).
			filename: "subjectCommonNamePrintableStringBadAlpha.pem",
			expected: lint.LintResult{
				Status:  lint.Error,
				Details: "RawSubject attr oid 2.5.4.3 encoded PrintableString contained illegal characters",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputPath := fmt.Sprintf("%s%s", util.TestCaseDir, tc.filename)
			result := lint.Lints["e_subject_printable_string_badalpha"].Execute(util.ReadCertificate(inputPath))
			if result.Status != tc.expected.Status {
				t.Errorf("expected result status %v was %v", tc.expected.Status, result.Status)
			}
			if result.Details != tc.expected.Details {
				t.Errorf("expected result details %q was %q", tc.expected.Details, result.Details)
			}
		})
	}
}
