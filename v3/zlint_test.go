package zlint

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/zmap/zlint/v3/util"

	"github.com/zmap/zcrypto/x509"
	"github.com/zmap/zlint/v3/lint"
)

func TestLintNames(t *testing.T) {
	allowedPrefixes := []string{
		"n_", // lints.Notice
		"w_", // lints.Warn
		"e_", // lints.Error
	}

	for _, name := range lint.GlobalRegistry().Names() {
		var valid bool
		for _, prefix := range allowedPrefixes {
			if strings.HasPrefix(name, prefix) {
				valid = true
				break
			}
		}
		if !valid {
			t.Errorf("lint name %q does not start with an allowed prefix (%v)\n",
				name, allowedPrefixes)
		}
	}
}

type configurableTestLint struct {
	A     string
	B     int
	C     map[string]string
	wantA string
	wantB int
	wantC map[string]string
}

func NewConfigurableTestLint() lint.LintInterface {
	return &configurableTestLint{C: make(map[string]string, 0), wantC: make(map[string]string, 0)}
}

func (l *configurableTestLint) Configure() interface{} {
	return l
}

func (l *configurableTestLint) CheckApplies(c *x509.Certificate) bool {
	return true
}

func (l *configurableTestLint) Execute(c *x509.Certificate) *lint.LintResult {
	if l.A != l.wantA {
		return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("A got %v, want %v", l.A, l.wantA)}
	}
	if l.B != l.wantB {
		return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("B got %v, want %v", l.B, l.wantB)}
	}
	if !reflect.DeepEqual(l.C, l.wantC) {
		return &lint.LintResult{Status: lint.Error, Details: fmt.Sprintf("C got %v, want %v", l.C, l.wantC)}
	}
	return &lint.LintResult{Status: lint.Pass}
}

func TestWithDefaultConfiguration(t *testing.T) {
	lint.RegisterLint(&lint.Lint{
		Name:          "library_usage_test_default_config",
		Description:   "CA Certificates subject field MUST not be empty and MUST have a non-empty distinguished name",
		Citation:      "RFC 5280: 4.1.2.6",
		Source:        lint.RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint:          NewConfigurableTestLint,
	})
	registry, err := lint.GlobalRegistry().Filter(lint.FilterOptions{
		IncludeNames: []string{"library_usage_test_default_config"},
	})
	if err != nil {
		t.Fatal(err)
	}
	got := LintCertificateEx(&x509.Certificate{
		NotAfter:  time.Now().Add(time.Hour),
		NotBefore: time.Now().Add(-time.Hour),
	}, registry)
	result, ok := got.Results["library_usage_test_default_config"]
	if !ok {
		t.Fatal("no results found, perhaps the lint never ran?")
	}
	if result.Status != lint.Pass {
		t.Fatalf("expected lint to pass, got %v (%s)", result.Status, result.Details)
	}
}

func TestWithNonDefaultConfiguration(t *testing.T) {
	lint.RegisterLint(&lint.Lint{
		Name:          "library_usage_test_non_default_config",
		Description:   "CA Certificates subject field MUST not be empty and MUST have a non-empty distinguished name",
		Citation:      "RFC 5280: 4.1.2.6",
		Source:        lint.RFC5280,
		EffectiveDate: util.RFC2459Date,
		Lint: func() lint.LintInterface {
			return &configurableTestLint{
				C:     make(map[string]string, 0),
				wantA: "the greatest song in the world",
				wantB: 42,
				wantC: map[string]string{
					"something": "else",
					"anything":  "at all",
				}}
		},
	})
	registry, err := lint.GlobalRegistry().Filter(lint.FilterOptions{
		IncludeNames: []string{"library_usage_test_non_default_config"},
	})
	if err != nil {
		t.Fatal(err)
	}
	config, err := lint.NewConfigFromString(`
[library_usage_test_non_default_config]
A = "the greatest song in the world"
B = 42

[library_usage_test_non_default_config.C]
something = "else"
anything = "at all"
`)
	if err != nil {
		t.Fatal(err)
	}
	registry.SetConfiguration(config)
	got := LintCertificateEx(&x509.Certificate{
		NotAfter:  time.Now().Add(time.Hour),
		NotBefore: time.Now().Add(-time.Hour),
	}, registry)
	result, ok := got.Results["library_usage_test_non_default_config"]
	if !ok {
		t.Fatal("no results found, perhaps the lint never ran?")
	}
	if result.Status != lint.Pass {
		t.Fatalf("expected lint to pass, got %v (%s)", result.Status, result.Details)
	}
}
