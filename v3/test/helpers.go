package test

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

// Contains resources necessary to the Unit Test Cases

import (
	"encoding/pem"
	"fmt"
	"os"
	"testing"

	"strings"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

// TestLint executes the given lintName against a certificate read from
// a testcert data file with the given filename. Filenames should be relative to
// `testdata/` and not absolute file paths.
//
// Important: TestLint is only appropriate for unit tests. It will panic if the
// lintName is not known or if the testCertFilename can not be loaded, or if the
// lint result is nil.
//
//nolint:revive
func TestLint(lintName string, testCertFilename string) *lint.LintResult {
	return TestLintWithConfig(lintName, testCertFilename, "")
}

func TestLintWithConfig(lintName string, testCertFilename string, configuration string) *lint.LintResult {
	config, err := lint.NewConfigFromString(configuration)
	if err != nil {
		panic(err)
	}
	return TestLintCert(lintName, ReadTestCert(testCertFilename), config)
}

// TestRevocationListLint executes the given lintName against a CRL read from
// a testcrl data file with the given filename. Filenames should be relative to
// `testdata/` and not absolute file paths.
//
//nolint:revive
func TestRevocationListLint(tb testing.TB, lintName string, testCRLFilename string) *lint.LintResult {
	tb.Helper()
	return TestRevocationListLintWithConfig(tb, lintName, testCRLFilename, "")
}

func TestRevocationListLintWithConfig(tb testing.TB, lintName string, testCRLFilename string, configuration string) *lint.LintResult {
	tb.Helper()
	config, err := lint.NewConfigFromString(configuration)
	if err != nil {
		tb.Fatal(err)
	}
	return TestLintRevocationList(tb, lintName, ReadTestRevocationList(tb, testCRLFilename), config)
}

// TestLintCert executes a lint with the given name against an already parsed
// certificate. This is useful when a unit test reads a certificate from disk
// and then mutates it in some way before trying to lint it.
//
// Important: TestLintCert is only appropriate for unit tests. It will panic if
// the lintName is not known or if the lint result is nil.
//
//nolint:revive
func TestLintCert(lintName string, cert *x509.Certificate, ctx lint.Configuration) *lint.LintResult {
	l := lint.GlobalRegistry().CertificateLints().ByName(lintName)
	if l == nil {
		panic(fmt.Sprintf(
			"Lint name %q does not exist in lint.Lints. "+
				"Did you forget to RegisterLint?\n",
			lintName))
	}
	res := l.Execute(cert, ctx)
	// We never expect a lint to return a nil LintResult
	if res == nil {
		panic(fmt.Sprintf(
			"Running lint %q on test certificate generated a nil LintResult.\n",
			lintName))
	}
	return res
}

// TestLintRevocationList executes a lint with the given name against an already parsed
// revocation list. This is useful when a unit test reads a revocation list from disk
// and then mutates it in some way before trying to lint it.
//
//nolint:revive
func TestLintRevocationList(tb testing.TB, lintName string, crl *x509.RevocationList, ctx lint.Configuration) *lint.LintResult {
	tb.Helper()
	l := lint.GlobalRegistry().RevocationListLints().ByName(lintName)
	if l == nil {
		tb.Fatalf(
			"Lint name %q does not exist in lint.Lints. "+
				"Did you forget to RegisterLint?\n",
			lintName)
	}
	res := l.Execute(crl, ctx)
	// We never expect a lint to return a nil LintResult
	if res == nil {
		tb.Fatalf(
			"Running lint %q on test revocation list generated a nil LintResult.\n",
			lintName)
	}
	return res
}

// ReadTestCert loads a x509.Certificate from the given inPath which is assumed
// to be relative to `testdata/`.
//
// Important: ReadTestCert is only appropriate for unit tests. It will panic if
// the inPath file can not be loaded.
func ReadTestCert(inPath string) *x509.Certificate {
	fullPath := fmt.Sprintf("../../testdata/%s", inPath)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		panic(fmt.Sprintf(
			"Unable to read test certificate from %q - %q "+
				"Does a unit test have an incorrect test file name?\n",
			fullPath, err))
	}

	if strings.Contains(string(data), "-BEGIN CERTIFICATE-") {
		block, _ := pem.Decode(data)
		if block == nil {
			panic(fmt.Sprintf(
				"Failed to PEM decode test certificate from %q - "+
					"Does a unit test have a buggy test cert file?\n",
				fullPath))
		}
		data = block.Bytes
	}

	theCert, err := x509.ParseCertificate(data)
	if err != nil {
		panic(fmt.Sprintf(
			"Failed to parse x509 test certificate from %q - %q "+
				"Does a unit test have a buggy test cert file?\n",
			fullPath, err))
	}

	return theCert
}

// ReadTestRevocationList loads a x509.RevocationList from the given inPath which is assumed
// to be relative to `testdata/`.
//
// Important: ReadTestRevocationList is only appropriate for unit tests. It will panic if
// the inPath file can not be loaded.
func ReadTestRevocationList(tb testing.TB, inPath string) *x509.RevocationList {
	tb.Helper()
	fullPath := fmt.Sprintf("../../testdata/%s", inPath)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		tb.Fatalf(
			"Unable to read test revocation list from %q - %q "+
				"Does a unit test have an incorrect test file name?\n",
			fullPath, err)
	}

	if strings.Contains(string(data), "-BEGIN X509 CRL-") {
		block, _ := pem.Decode(data)
		if block == nil { //nolint: staticcheck // tb.Fatalf exits
			tb.Fatalf(
				"Failed to PEM decode test revocation list from %q - "+
					"Does a unit test have a buggy test cert file?\n",
				fullPath)
		}
		data = block.Bytes //nolint: staticcheck // tb.Fatalf exits
	}

	theCrl, err := x509.ParseRevocationList(data)
	if err != nil {
		tb.Fatalf(
			"Failed to parse x509 test certificate from %q - %q "+
				"Does a unit test have a buggy test cert file?\n",
			fullPath, err)
	}

	return theCrl
}
