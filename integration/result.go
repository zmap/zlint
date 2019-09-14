// +build integration

package integration

import "fmt"

type result struct {
	FatalCount  uint8 `json:",omitempty"`
	ErrCount    uint8 `json:",omitempty"`
	WarnCount   uint8 `json:",omitempty"`
	NoticeCount uint8 `json:",omitempty"`
}

func (r result) fullPass() bool {
	return r.FatalCount == 0 && r.ErrCount == 0 && r.WarnCount == 0 && r.NoticeCount == 0
}

func (r result) String() string {
	return fmt.Sprintf("fatals: %d\terrs: %d\twarns: %d\tinfos: %d",
		r.FatalCount, r.ErrCount, r.WarnCount, r.NoticeCount)
}

// certResult combines a Result with a Fingerprint.
type certResult struct {
	Fingerprint string
	Result      result
}

func (cr certResult) String() string {
	return fmt.Sprintf("%q\t%s", cr.Fingerprint, cr.Result)
}
