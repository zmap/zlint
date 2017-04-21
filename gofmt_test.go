package main

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestGofmt(t *testing.T) {
	cmd := exec.Command("/bin/sh", "-c", "gofmt -l main.go")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	if out.String() != "" {
		t.Error(
			"main.go not gofmt'ed",
		)
	}
	out.Reset()
	cmd = exec.Command("/bin/sh", "-c", "gofmt -l lints/*")
	cmd.Stdout = &out
	cmd.Run()
	if out.String() != "" {
		t.Error(
			"lints not gofmt'ed",
		)
	}
	out.Reset()
	cmd = exec.Command("/bin/sh", "-c", "gofmt -l util/*")
	cmd.Stdout = &out
	cmd.Run()
	if out.String() != "" {
		t.Error(
			"util not gofmt'ed",
		)
	}
}
