package test

/*
 * ZLint Copyright 2023 Regents of the University of Michigan
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

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"testing"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/util"

	"github.com/zmap/zlint/v3/lint"
)

func init() {
	// This is a complication caused https://github.com/zmap/zlint/issues/696
	//
	// This test package required access to the test certificate directory, however
	// the ReadTestCert testing helper function assumes that your PWD is one of the
	// lint genre directories.
	//
	// ReadTestCert was changed to operate from the root of the repo to accommodate this
	// test package, however that broke downstream consumers who were dependent on the
	// relative path building behavior.
	err := os.Chdir("../lints/rfc")
	if err != nil {
		panic(err)
	}
}

type caCommonNameMissing struct {
	BeerHall string
	Working  *lint.CABFBaselineRequirementsConfig
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "e_ca_common_name_missing2",
		Description:   "CA Certificates common name MUST be included.",
		Citation:      "BRs: 7.1.4.3.1",
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.CABV148Date,
		Lint:          NewCaCommonNameMissing,
	})
}

func (l *caCommonNameMissing) Configure() interface{} {
	return l
}

func NewCaCommonNameMissing() lint.LintInterface {
	return &caCommonNameMissing{}
}

func (l *caCommonNameMissing) CheckApplies(c *x509.Certificate) bool {
	return util.IsCACert(c)
}

func (l *caCommonNameMissing) Execute(c *x509.Certificate) *lint.LintResult {
	if c.Subject.CommonName == "" {
		return &lint.LintResult{Status: lint.Error, Details: l.BeerHall}
	} else {
		return &lint.LintResult{Status: lint.Pass, Details: l.BeerHall}
	}
}

func TestCaCommonNameMissing(t *testing.T) {
	inputPath := "caCommonNameMissing.pem"
	expected := lint.Error
	out := TestLint("e_ca_common_name_missing2", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaCommonNameNotMissing(t *testing.T) {
	inputPath := "caCommonNameNotMissing.pem"
	expected := lint.Pass
	out := TestLint("e_ca_common_name_missing2", inputPath)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaCommonNameNotMissing2(t *testing.T) {
	inputPath := "caCommonNameNotMissing.pem"
	expected := lint.Pass
	config := `
[e_ca_common_name_missing2]
BeerHall = "liedershousen"
`
	out := TestLintWithConfig("e_ca_common_name_missing2", inputPath, config)
	if out.Details != "liedershousen" {
		t.Fatalf("unexpected output details, got '%s' want %s", out.Details, "liedershousen")
	}
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

func TestCaCommonNameNotMissing3(t *testing.T) {
	inputPath := "caCommonNameNotMissing.pem"
	expected := lint.Pass
	config := `
[e_ca_common_name_missing2]
BeerHall = "liedershousenssss"
`
	out := TestLintWithConfig("e_ca_common_name_missing2", inputPath, config)
	if out.Details != "liedershousenssss" {
		t.Fatalf("unexpected output details, got '%s' want %s", out.Details, "liedershousenssss")
	}
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
}

// This exercises the thread safety our configurable lints. This is because
// the lints use to be global singletons before we swapped them over to
// running as single instances. However, it is a good exercise to keep around.
func TestConcurrency(t *testing.T) {
	inputPath := "caCommonNameNotMissing.pem"
	expected := lint.Pass
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			num := strconv.Itoa(rand.Intn(9999))
			config := fmt.Sprintf(`
[e_ca_common_name_missing2]
BeerHall = "%s"
`, num)
			out := TestLintWithConfig("e_ca_common_name_missing2", inputPath, config)
			if out.Details != num {
				t.Errorf("wanted %s got %s", num, out.Details)
			}
			if out.Status != expected {
				t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
			}
		}()
	}
	wg.Wait()
}

func TestCaCommonNameNotMissing4(t *testing.T) {
	inputPath := "caCommonNameNotMissing.pem"
	expected := lint.Pass
	config := `
[CABF_BR]
DoesItWork = "yes, yes it does"

[e_ca_common_name_missing2]
BeerHall = "liedershousenssss"
`
	out := TestLintWithConfig("e_ca_common_name_missing2", inputPath, config)
	if out.Status != expected {
		t.Errorf("%s: expected %s, got %s", inputPath, expected, out.Status)
	}
	if out.Details != "liedershousenssss" {
		t.Fatalf("unexpected output details, got '%s' want %s", out.Details, "liedershousenssss")
	}
}

type LintEmbedsAConfiguration struct {
	configuration                        embeddedConfiguration
	SomeOtherFieldThatWeDontWantToExpose int
}

type embeddedConfiguration struct {
	IsWebPKI bool `comment:"Indicates that the certificate is intended for the Web PKI." toml:"is_web_pki"`
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:          "w_web_pki_cert",
		Description:   "CA Certificates SHOULD....something....about the web pki",
		Citation:      "BRs: 7.1.4.3.1",
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.CABV148Date,
		Lint:          NewLintEmbedsAConfiguration,
	})
}

// A pointer to an embedded struct may be passed to the framework
// if the author does not wish to expose certain fields in their primary struct.
func (l *LintEmbedsAConfiguration) Configure() interface{} {
	return &l.configuration
}

func NewLintEmbedsAConfiguration() lint.LintInterface {
	return &LintEmbedsAConfiguration{configuration: embeddedConfiguration{}}
}

func (l *LintEmbedsAConfiguration) CheckApplies(c *x509.Certificate) bool {
	return util.IsCACert(c)
}

func (l *LintEmbedsAConfiguration) Execute(c *x509.Certificate) *lint.LintResult {
	if l.configuration.IsWebPKI {
		return &lint.LintResult{Status: lint.Warn, Details: "Time for a beer run!"}
	} else {
		return &lint.LintResult{Status: lint.Pass}
	}
}
