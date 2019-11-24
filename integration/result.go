// +build integration

package integration

import (
	"fmt"

	"github.com/zmap/zlint/lints"
)

type resultCount struct {
	FatalCount  uint8 `json:",omitempty"`
	ErrCount    uint8 `json:",omitempty"`
	WarnCount   uint8 `json:",omitempty"`
	NoticeCount uint8 `json:",omitempty"`
}

// TODO(@cpu): Accept a threshold argument so that (for e.g. notices could be
// counted as passing)
func (r resultCount) fullPass() bool {
	return r.FatalCount == 0 && r.ErrCount == 0 && r.WarnCount == 0 && r.NoticeCount == 0
}

func (r resultCount) String() string {
	return fmt.Sprintf("fatals: %-4d errs: %-4d warns: %-4d infos: %-4d",
		r.FatalCount, r.ErrCount, r.WarnCount, r.NoticeCount)
}

// Inc increases the resultCount count for the given lint status level.
func (r *resultCount) Inc(status lints.LintStatus) {
	switch status {
	case lints.Notice:
		r.NoticeCount++
	case lints.Warn:
		r.WarnCount++
	case lints.Error:
		r.ErrCount++
	case lints.Fatal:
		r.FatalCount++
	}
}

// certResult combines a Result (overall count of lint results by type) with
// a LintSummary (map from lint name to a Notice/Warn/Error/Fatal result) for
// a specific cert Fingerprint.
type certResult struct {
	Fingerprint string
	Result      resultCount
	LintSummary map[string]lints.LintStatus
}

func (cr certResult) String() string {
	return fmt.Sprintf("%q\t%s", cr.Fingerprint, cr.Result)
}
