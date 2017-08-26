package zlint

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestGofmt(t *testing.T) {
	globs := []string{
		"*.go",
		"cmd/*.go",
		"lints/*.go",
		"util/*.go",
	}
	for _, glob := range globs {
		gofmtCmd := "gofmt -s -l " + glob
		cmd := exec.Command("/bin/sh", "-c", gofmtCmd)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		if out.String() != "" {
			t.Errorf("glob %s not gofmt'ed", glob)
		}
	}
}
